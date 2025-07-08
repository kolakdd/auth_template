// Package models
package models

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUserDtoRes struct {
	GUID      uuid.UUID `json:"guid" example:" a81bc81b-dead-4e5d-abff-90865d1e13b1"`
	Name      string    `json:"name" example:"Nikolay"`
	CreatedAt time.Time `json:"createdAt" example:"2025-07-07T07:01:27.73104763Z"`
}

type User struct {
	GUID        uuid.UUID `json:"guid" example:" a81bc81b-dead-4e5d-abff-90865d1e13b1"`
	Name        string    `json:"name" example:"Nikolay"`
	Deactivated bool      `json:"deactivated" example:"false"`
	LastLoginAt time.Time `json:"lastLoginAt" example:"2025-07-07T07:01:27.73104763Z"`
	CreatedAt   time.Time `json:"createdAt" example:"2025-07-07T07:01:27.73104763Z"`
}

func NewUserDB(name string) User {
	timeNow := time.Now().UTC()
	return User{
		GUID:        uuid.New(),
		Name:        name,
		Deactivated: false,
		LastLoginAt: timeNow,
		CreatedAt:   timeNow,
	}
}

type RegisterUserDtoReq struct {
	Name string `json:"name" example:"Nikolay"`
}
