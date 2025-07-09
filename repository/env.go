package repository

import (
	"fmt"
	"os"
	"strconv"
)

type RepositoryEnv interface {
	GetDatabaseDSN() string
	GetAPIMode() string
	GetRefreshTokenExpiredSec() int
	GetSecret() string
}

type repositoryEnv struct {
	pgUser     string `env:"POSTGRES_USER"`
	pgPassword string `env:"POSTGRES_PASSWORD"`
	pgHost     string `env:"PG_HOST"`
	dbName     string `env:"DB_NAME"`
	dbPort     int    `env:"DB_PORT"`

	apiSecret string `env:"API_SECRET"`

	accessTokenExpired  int `env:"ACCESS_TOKEN_EXPIRED"`
	refreshTokenExpired int `env:"REFRESH_TOKEN_EXPIRED"`

	mode string `env:"MODE"`
}

func NewRepoEnv() RepositoryEnv {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgHost := os.Getenv("PG_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := parseInt("DB_PORT")

	apiSecret := os.Getenv("API_SECRET")

	accessTokenExpired := parseInt("ACCESS_TOKEN_EXPIRED")
	refreshTokenExpired := parseInt("REFRESH_TOKEN_EXPIRED")

	mode := os.Getenv("MODE")

	return &repositoryEnv{pgUser, pgPassword, pgHost, dbName, dbPort, apiSecret, accessTokenExpired, refreshTokenExpired, mode}
}

func parseInt(evnKey string) int {
	v := os.Getenv(evnKey)
	vInt, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Sprintf("err while parse env key=%s", evnKey))
	}
	return vInt

}
func (r *repositoryEnv) GetDatabaseDSN() string {
	user := r.pgUser
	password := r.pgPassword
	host := r.pgHost
	db := r.dbName
	port := r.dbPort

	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, db, port)
}

func (r *repositoryEnv) GetAPIMode() string {
	return r.mode
}

func (r *repositoryEnv) GetRefreshTokenExpiredSec() int {
	return r.refreshTokenExpired
}
func (r *repositoryEnv) GetSecret() string {
	return r.apiSecret
}
