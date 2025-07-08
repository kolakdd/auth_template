// Package middleware
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kolakdd/auth_template/service"
)

// AuthMiddleware проверяет:
// 1) наличие хедера Authorization: Bearer {token}
// 2) валидирует токен по времени протухания и по структуре
// 3) наличие пользователя по sub
// 4) наличие деактивировации на аккаунте
func AuthMiddleware(authService service.ServiceAuthI, userService service.ServiceUserI) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return userService.AuthMiddlewareFunc(c)
	}
}
