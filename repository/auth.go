package repository

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"github.com/kolakdd/auth_template/models"
	"github.com/kolakdd/auth_template/query"
	"github.com/kolakdd/auth_template/secure"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RepositoryAuth interface {
	CreateRefreshToken(token string, userGUID uuid.UUID) (*models.RefreshToken, error)
	CreateInvalidAccessToken(token uuid.UUID, userGUID uuid.UUID) (*models.InvalidAccessToken, error)
	GetInvalidAccessToken(guid uuid.UUID) (bool, *models.InvalidAccessToken)
	GetRefreshToken(token string, userGUID uuid.UUID) (*models.RefreshToken, error)
	DeleteRefreshToken(tokenHash string) error
	ValidateAuthHeader(secret, authHeader string) (*secure.AccessToken, error)
}

type repositoryAuth struct {
	db *gorm.DB
}

func NewRepoAuth(db *gorm.DB) RepositoryAuth {
	return &repositoryAuth{
		db: db,
	}
}

func (r *repositoryAuth) CreateRefreshToken(token string, userGUID uuid.UUID) (*models.RefreshToken, error) {
	hashToken := secure.HashRefreshToken(token)
	q := query.Use(r.db)
	refreshToken := models.RefreshTokenDBNew(hashToken, userGUID)
	if err := q.RefreshToken.Create(&refreshToken); err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *repositoryAuth) CreateInvalidAccessToken(guid uuid.UUID, userGUID uuid.UUID) (*models.InvalidAccessToken, error) {
	q := query.Use(r.db)
	invalidAccessToken := models.InvalidAccessTokenDBNew(guid, userGUID)
	if err := q.InvalidAccessToken.Create(&invalidAccessToken); err != nil {
		return nil, err
	}
	return &invalidAccessToken, nil
}

func (r *repositoryAuth) GetInvalidAccessToken(guid uuid.UUID) (bool, *models.InvalidAccessToken) {
	var tokens []models.InvalidAccessToken
	if err := r.db.Where("guid = ?", guid).Find(&tokens).Error; err != nil {
		return false, nil
	}
	if len(tokens) == 0 {
		return false, nil
	}
	return true, &tokens[0]
}

// GetRefreshToken получает рефреш токены пользователя и сравнивает их bcrypt в поиске подходящего
func (r *repositoryAuth) GetRefreshToken(token string, userGUID uuid.UUID) (*models.RefreshToken, error) {
	if len(token) > 72 {
		token = token[:72]
	}
	var tokens []models.RefreshToken

	if err := r.db.Where("user_guid = ?", userGUID).Find(&tokens).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(tokens); i++ {
		err := bcrypt.CompareHashAndPassword([]byte(tokens[i].TokenHash), []byte(token))
		if err == nil {
			return &tokens[i], nil
		}
	}
	return nil, fmt.Errorf("token not found")
}

func (r *repositoryAuth) DeleteRefreshToken(hashToken string) error {
	refreshTokenDB := new(models.RefreshToken)
	return r.db.Where("token_hash = ?", hashToken).Delete(&refreshTokenDB).Error
}

// ValidateAuthHeader проверят access токен на правильность формата
func (r *repositoryAuth) ValidateAuthHeader(secret, authHeader string) (*secure.AccessToken, error) {
	accessToken, err := secure.ValidateAccessToken(secret, strings.Split(authHeader, " ")[1])
	if err != nil {
		slog.Warn("Validate token ", "err ", err)
		return nil, fmt.Errorf("access token not valid")
	}
	return accessToken, nil
}
