package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/kolakdd/auth_template/database"
	_ "github.com/kolakdd/auth_template/docs"
	"github.com/kolakdd/auth_template/repository"
	"github.com/kolakdd/auth_template/routes"
	gormlogger "gorm.io/gorm/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	envRepo := repository.NewRepoEnv()

	db, err := database.Connect(envRepo.GetDatabaseDSN())
	if err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	app := fiber.New()

	if envRepo.GetAPIMode() == "DEBUG" {
		db.Logger = gormlogger.Default.LogMode(gormlogger.Info)
	} else {
		db.Logger = gormlogger.Default.LogMode(gormlogger.Error)
	}

	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	routes.New(app, db, envRepo)
	log.Fatal(app.Listen(":3000"))
}
