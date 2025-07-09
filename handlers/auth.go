// Package handlers
package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kolakdd/auth_template/httputil"
	"github.com/kolakdd/auth_template/models"
	"github.com/kolakdd/auth_template/service"
)

type AuthHandler struct {
	sa service.ServiceAuthI
	su service.ServiceUserI
}

func NewAuthHandler(sa service.ServiceAuthI, su service.ServiceUserI) *AuthHandler {
	return &AuthHandler{sa, su}
}

// RegisterUser registers a new user
// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя в системе с заданным именем. Возвращает информацию о пользователе с GUID
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.RegisterUserDtoReq true "User data"
// @Success 200 {object} httputil.ResponseHTTP{data=models.User} "Success response"
// @Failure 400 {object} httputil.ResponseHTTP "Bad Request"
// @Router /api/v1/register [post]
func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	userDto := new(models.RegisterUserDtoReq)
	if err := c.BodyParser(userDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest(err.Error()))
	}
	user, err := h.su.RegisterUser(userDto)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest(err.Error()))
	}
	return c.JSON(httputil.ResponseHTTP{
		Success: true,
		Message: "User registered successfully",
		Data:    user,
	})
}

// LoginGUID log in user in system, provide access and refresh tokens
// @Summary Логин в систему
// @Description Генерирует access и refresh токен по GUID пользователя для авторизации в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param        guid    path     string  true  "Guid авторизируемого пользователя"
// @Success 200 {object} httputil.ResponseHTTP{data=models.LoginTokens}
// @Failure 400 {object} httputil.ResponseHTTP "Bad Request"
// @Router /api/v1/login/{guid} [post]
func (h *AuthHandler) LoginGUID(c *fiber.Ctx) error {
	guid := c.Params("guid")
	userGUID, err := uuid.Parse(guid)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest("param guid not exist"))
	}
	tokens, err := h.sa.LoginUser(userGUID, c.IP(), string(c.Context().UserAgent()))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest(err.Error()))
	}
	return c.JSON(httputil.ResponseHTTP{
		Success: true,
		Message: "Login success.",
		Data:    tokens,
	})
}

// RefreshTokens refresh access and refresh tokens
// @Summary Обновление токенов пользователя
// @Description Генерирует access и refresh токен по ранее предоставленным при авторизации токеном
// @Tags auth
// @Accept json
// @Produce json
// @Param tokens body models.LoginTokens true "Tokes"
// @Success 200 {object} httputil.ResponseHTTP{data=models.LoginTokens}
// @Failure 400 {object} httputil.ResponseHTTP "Bad Request"
// @Failure 500 {object} httputil.ResponseHTTP "Internal Server Error"
// @Router /api/v1/refresh [post]
func (h *AuthHandler) RefreshTokens(c *fiber.Ctx) error {
	tokensDto := new(models.LoginTokens)
	if err := c.BodyParser(tokensDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest(err.Error()))
	}

	res, err := h.sa.RefreshToken(tokensDto, c.IP(), string(c.Context().UserAgent()))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(httputil.BadRequest(err.Error()))

	}
	return c.JSON(httputil.ResponseHTTP{
		Success: true,
		Message: "Refresh success.",
		Data:    res,
	})
}
