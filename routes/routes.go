// Package routes
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/kolakdd/auth_template/handlers"
	"github.com/kolakdd/auth_template/middleware"
	"github.com/kolakdd/auth_template/repository"
	"github.com/kolakdd/auth_template/service"
	"gorm.io/gorm"
)

// New create an instance of Book app routes
func New(app *fiber.App, db *gorm.DB, envRepo repository.RepositoryEnv) error {
	app.Get("/swagger/*", swagger.HandlerDefault)

	userRepo := repository.NewRepoUser(db)
	authRepo := repository.NewRepoAuth(db)
	authService := service.NewServiceAuth(authRepo, userRepo, envRepo)
	userService := service.NewServiceUser(authRepo, userRepo, envRepo)

	authMiddleware := middleware.AuthMiddleware(authService, userService)

	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{
			"message": "v1",
		})
		return c.Next()
	})

	v1.Post("/webhook", handlers.WebhookExample)

	authHandler := handlers.NewAuthHandler(authService, userService)

	v1.Post("/register", authHandler.RegisterUser)
	v1.Post("/login/:guid", authHandler.LoginGUID)
	v1.Post("/refresh", authHandler.RefreshTokens)

	userHandler := handlers.NewUserHandler(userService)

	user := v1.Group("/user")
	user.Use(authMiddleware)

	user.Get("/me", userHandler.UserMe)
	user.Get("/unlogin", userHandler.UnloginMe)
	user.Get("/deactivate", userHandler.DeactivateMe)

	return nil
}
