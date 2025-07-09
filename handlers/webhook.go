package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kolakdd/auth_template/httputil"
	"github.com/kolakdd/auth_template/models"
)

// WebhookExample provide test route for webhook action
// @Summary Тестовый веб-хук который позволяет проверить работы колбека
// @Description Информирует о правильной работе запроса к вуб-хуку
// @Tags webhook
// @Accept json
// @Produce json
// @Param        Authorization	  header    string    true   	"Заголовок авторизации. Пример: Bearer {token}"
// @Success 200 {object} httputil.ResponseHTTP{}
// @Router /api/v1/webhook [get]
func WebhookExample(c *fiber.Ctx) error {
	dto := new(models.WebHookDto)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest(err.Error()))
	}
	fmt.Println(" << WebHook get message", dto)
	return c.JSON(httputil.ResponseHTTP{
		Success: true,
		Message: "<< WebHook get message.",
		Data:    nil,
	})
}
