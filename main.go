package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/kolakdd/auth_template/database"
	_ "github.com/kolakdd/auth_template/docs"
	"github.com/kolakdd/auth_template/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	db, err := database.Connect()
	if err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	routes.New(app, db)
	log.Fatal(app.Listen(":3000"))
}
