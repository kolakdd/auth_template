package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kolakdd/auth_template/httputil"
	"github.com/kolakdd/auth_template/models"
	"github.com/kolakdd/auth_template/repository"
	"github.com/kolakdd/auth_template/secure"
)

type ServiceUserI interface {
	AuthMiddlewareFunc(c *fiber.Ctx) error
	RegisterUser(dto *models.RegisterUserDtoReq) (*models.User, error)
	UserMe(c *fiber.Ctx) *models.User
	DeativateMe(authHeader string) (*models.User, error)
	UnloginMe(c *fiber.Ctx) (*models.InvalidAccessToken, error)
}

type ServiceUser struct {
	RAuth repository.RepositoryAuth
	RUser repository.RepositoryUser
	rEnv  repository.RepositoryEnv
}

func NewServiceUser(rAuth repository.RepositoryAuth, rUser repository.RepositoryUser, rEnv repository.RepositoryEnv) ServiceUserI {
	return &ServiceUser{rAuth, rUser, rEnv}
}

func (s *ServiceUser) RegisterUser(dto *models.RegisterUserDtoReq) (*models.User, error) {
	user, err := s.RUser.NewUser(dto.Name)
	if err != nil {
		return nil, fmt.Errorf("error while create user")
	}
	return user, nil
}

func (s *ServiceUser) UserMe(c *fiber.Ctx) *models.User {
	user := c.Locals("user").(*models.User)
	return user
}

// DeativateMe деактивирует пользователя, пользователь больше не сможет взаимодействовать с аккаунтом
func (s *ServiceUser) DeativateMe(authHeader string) (*models.User, error) {
	accessToken, err := s.RAuth.ValidateAuthHeader(s.rEnv.GetSecret(), authHeader)
	if err != nil {
		return nil, err
	}
	user, err := s.RUser.GetUser(accessToken.Sub)
	if err != nil {
		return nil, err
	}
	if user.Deactivated {
		return nil, fmt.Errorf("user deactivated")
	}

	user, err = s.RUser.DeactivateUser(accessToken.Sub)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UnloginMe деавторизирует пользователя путем добавления id токена в черный лист
func (s *ServiceUser) UnloginMe(c *fiber.Ctx) (*models.InvalidAccessToken, error) {
	user := c.Locals("user").(*models.User)
	accessToken := c.Locals("accessToken").(*secure.AccessToken)

	ivalidToken, err := s.RAuth.CreateInvalidAccessToken(accessToken.ID, user.GUID)
	if err != nil {
		return nil, err
	}
	return ivalidToken, nil
}

func (s *ServiceUser) AuthMiddlewareFunc(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest("Invalid headers muts be \"Authorization\": \"Bearer {token}\""))
	}
	accessToken, err := s.RAuth.ValidateAuthHeader(s.rEnv.GetSecret(), authHeader)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid access token",
		})
	}

	exist, _ := s.RAuth.GetInvalidAccessToken(accessToken.ID)
	if exist {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "access token deactivated",
		})
	}
	user, err := s.RUser.GetUser(accessToken.Sub)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	if user.Deactivated {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user deactivated",
		})
	}

	c.Locals("user", user)               // *models.User
	c.Locals("accessToken", accessToken) // *secure.AccessToken
	return c.Next()
}
