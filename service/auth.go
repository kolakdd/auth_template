// Package service
package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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
	rEnv  repository.RepositoryEnv
}

func NewServiceAuth(rAuth repository.RepositoryAuth, rUser repository.RepositoryUser, rEnv repository.RepositoryEnv) ServiceAuthI {
	return &ServiceAuth{rAuth, rUser, rEnv}
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
	validatedAccessToken, err := secure.DecodeAccessToken(s.rEnv.GetSecret(), dto.AccessToken)
	if err != nil {
		return nil, err
	}
	exist, _ := s.rAuth.GetInvalidAccessToken(validatedAccessToken.ID)
	if exist {
		return nil, fmt.Errorf("access token deactivated")
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
		weebHookDto := models.WebHookDto{user.GUID, decodedRefreshToken.IP, ip, time.Now().UTC()}
		jsonDto, err := json.Marshal(weebHookDto)
		if err != nil {
			return nil, fmt.Errorf("rrror marshaling ")
		}
		http.Post("http://localhost:3000/api/v1/webhook", "application/json", bytes.NewBuffer(jsonDto))
	}

	if !strings.HasPrefix(userAgent, decodedRefreshToken.UserAgent) {
		_, err := s.rAuth.CreateInvalidAccessToken(validatedAccessToken.ID, user.GUID)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("new userAgent detected")
	}

	dif := time.Now().UTC().Sub(oldRefTokenDB.CreatedAt)
	if dif >= time.Duration(time.Duration(s.rEnv.GetRefreshTokenExpiredSec())*time.Second) {
		err = s.rAuth.DeleteRefreshToken(oldRefTokenDB.TokenHash)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("token expired")
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
