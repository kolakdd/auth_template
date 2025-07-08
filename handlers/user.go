package handlers

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kolakdd/auth_template/httputil"
	"github.com/kolakdd/auth_template/service"
)

type UserHandler struct {
	s service.ServiceUserI
}

func NewUserHandler(s service.ServiceUserI) *UserHandler {
	return &UserHandler{s: s}
}

// UserMe get info about user, by access token
// @Summary Получение GUID текущего пользователя и информации о нем."
// @Description По токену доступа возвращает информацию о пользователе. Необходим заголовок "Authorization": "Bearer {token}
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param        Authorization	  header    string    true   	"Заголовок авторизации. Пример: Bearer {token}"
// @Success 200 {object} httputil.ResponseHTTP{data=models.LoginTokens}
// @Failure 400 {object} httputil.ResponseHTTP "Bad Request"
// @Router /api/v1/user/me [get]
func (h *UserHandler) UserMe(c *fiber.Ctx) error {
	userMe := h.s.UserMe(c)
	return c.JSON(httputil.ResponseHTTP{
		Success: true,
		Message: "User me",
		Data:    userMe,
	})
}

// UnloginMe deauthorization user
// @Summary Деавторизация токена пользователя
// @Description По токену авторизации деактивирует пользователя так, что он не может запросить user/me или /refresh. Необходим заголовок "Authorization": "Bearer {token}
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param        Authorization	  header    string    true   	"Заголовок авторизации. Пример: Bearer {token}"
// @Success 200 {object} httputil.ResponseHTTP{data=models.User}
// @Failure 400 {object} httputil.ResponseHTTP "Invalid headers muts be "Authorization": "Bearer {token}""
// @Failure 403 {object} httputil.ResponseHTTP "Access denied"
// @Router /api/v1/user/unlogin [get]
func (h *UserHandler) UnloginMe(c *fiber.Ctx) error {
	userDeactivated, err := h.s.UnloginMe(c)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(httputil.BadRequest("Access denied"))
	}
	return c.JSON(httputil.ResponseHTTP{
		Success: true,
		Message: "User unlogin success.",
		Data:    userDeactivated,
	})
}

// DeactivateMe deactivate authorizatied user
// @Summary Диактивация авторизованного пользователя
// @Description Блокирует пользователя пользователя по токенгу авторизации. Необходим заголовок "Authorization": "Bearer {token}
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param        Authorization	  header    string    true   	"Заголовок авторизации. Пример: Bearer {token}"
// @Success 200 {object} httputil.ResponseHTTP{data=models.User}
// @Failure 400 {object} httputil.ResponseHTTP "Bad Request"
// @Failure 401 {object} httputil.ResponseHTTP "Unauthorized"
// @Router /api/v1/user/deactivate [get]
// @Deprecated
func (h *UserHandler) DeactivateMe(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest("Invalid headers muts be \"Authorization\": \"Bearer {token}\""))
	}

	userDeactivated, err := h.s.DeativateMe(authHeader)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(httputil.BadRequest("Access denied"))
	}
	return c.JSON(httputil.ResponseHTTP{
		Success: true,
		Message: "User deactivated.",
		Data:    userDeactivated,
	})
}
