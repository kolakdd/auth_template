package handlers

// import (
// 	"log/slog"
// 	"net/http"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/kolakdd/auth_template/database"
// 	"github.com/kolakdd/auth_template/httputil"
// 	"github.com/kolakdd/auth_template/models"
// 	"github.com/kolakdd/auth_template/secure"
// )

// // WebhookExample provide test route for webhook action
// // @Summary Тестовый веб-хук который позволяет проверить работы колбека."
// // @Description Инофрмаирует о правильной работе запроса к вуб-хуку
// // @Tags auth
// // @Accept json
// // @Produce json
// // @Security Bearer
// // @Param        Authorization	  header    string    true   	"Заголовок авторизации. Пример: Bearer {token}"
// // @Success 200 {object} httputil.ResponseHTTP{data=models.LoginTokens}
// // @Failure 400 {object} httputil.ResponseHTTP "Bad Request"
// // @Router /api/v1/auth/webhook [post]
// func WebhookExample(c *fiber.Ctx) error {
// 	authHeader := c.Get("Authorization")
// 	if !strings.HasPrefix(authHeader, "Bearer ") {
// 		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest("Invalid headers muts be \"Authorization\": \"Bearer {token}\""))
// 	}

// 	userGUID, err := secure.ValidateAccessToken(strings.Split(authHeader, " ")[1])
// 	if err != nil {
// 		slog.Warn("Validate token ", "err ", err)
// 		return c.Status(http.StatusUnauthorized).JSON(httputil.BadRequest("Access token not valid"))
// 	}

// 	user := new(models.User)
// 	err = database.DBConn.Where("GUID = ?", userGUID).Find(&user).Error
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(httputil.BadRequest("user not found "))
// 	}
// 	if user.Deactivated {
// 		return c.Status(http.StatusForbidden).JSON(httputil.BadRequest("user deactivated"))
// 	}

// 	return c.JSON(httputil.ResponseHTTP{
// 		Success: true,
// 		Message: "user me.",
// 		Data:    user,
// 	})
// }
