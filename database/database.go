// Package database
package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kolakdd/auth_template/models"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	db := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, db, portInt)
	DBConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := DBConn.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	initGormModels(DBConn)

	return DBConn, nil
}

// initGormModels apply db models and config gorm codegen
func initGormModels(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)
	g.ApplyBasic(
		models.User{},
		models.RefreshToken{},
		models.InvalidAccessToken{},
	)
	g.Execute()
}
