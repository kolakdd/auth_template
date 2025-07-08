// Package service
package service

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/kolakdd/auth_template/models"
	"github.com/kolakdd/auth_template/repository"
	"github.com/kolakdd/auth_template/secure"
)

type ServiceAuthI interface {
	LoginUser(id uuid.UUID, ip string, userAgent string) (*models.LoginTokens, error)
	RefreshToken(dto *models.LoginTokens, ip string, userAgent string) (*models.LoginTokens, error)
}

type ServiceAuth struct {
	rAuth repository.RepositoryAuth
	rUser repository.RepositoryUser
}

func NewServiceAuth(rAuth repository.RepositoryAuth, rUser repository.RepositoryUser) ServiceAuthI {
	return &ServiceAuth{rAuth, rUser}
}

func (s *ServiceAuth) LoginUser(userID uuid.UUID, ip string, userAgent string) (*models.LoginTokens, error) {
	user, err := s.rUser.GetUser(userID)
	if err != nil {
		return nil, err
	}
	if user.Deactivated {
		return nil, fmt.Errorf("user deactivated")
	}
	refTokenID := uuid.New()
	tokenAccess := secure.GenerateAccessToken(userID, refTokenID)
	tokenRefresh := secure.GenerateRefreshToken(refTokenID.String(), ip, userAgent)

	_, err = s.rAuth.CreateRefreshToken(tokenRefresh, user.GUID)
	if err != nil {
		return nil, err
	}
	return &models.LoginTokens{AccessToken: tokenAccess, RefreshToken: tokenRefresh}, nil
}

func (s *ServiceAuth) RefreshToken(dto *models.LoginTokens, ip string, userAgent string) (*models.LoginTokens, error) {
	validatedAccessToken, err := secure.DecodeAccessToken(dto.AccessToken)
	if err != nil {
		return nil, err
	}
	user, err := s.rUser.GetUser(validatedAccessToken.Sub)
	if err != nil {
		return nil, err
	}
	oldRefTokenDB, err := s.rAuth.GetRefreshToken(dto.RefreshToken, user.GUID)
	if err != nil {
		return nil, err
	}

	decodedRefreshToken, err := secure.DecodeRefreshToken(dto.RefreshToken)
	if err != nil {
		return nil, err
	}

	if user.Deactivated {
		return nil, fmt.Errorf("user deactivated")
	}

	if validatedAccessToken.Ref != decodedRefreshToken.ID {
		return nil, fmt.Errorf("bad token pair")
	}

	if decodedRefreshToken.IP != ip {
		// todo: send web hook
		fmt.Println("send web hook")
	}

	if !strings.HasPrefix(userAgent, decodedRefreshToken.UserAgent) {
		// todo: deauthorization tokens
		return nil, fmt.Errorf("new userAgent detected")
	}

	refTokenID := uuid.New()
	tokenAccess := secure.GenerateAccessToken(user.GUID, refTokenID)
	tokenRefresh := secure.GenerateRefreshToken(refTokenID.String(), ip, userAgent)

	err = s.rAuth.DeleteRefreshToken(oldRefTokenDB.TokenHash)
	if err != nil {
		return nil, err
	}
	_, err = s.rAuth.CreateRefreshToken(tokenRefresh, user.GUID)
	if err != nil {
		return nil, err
	}
	return &models.LoginTokens{AccessToken: tokenAccess, RefreshToken: tokenRefresh}, nil
}
