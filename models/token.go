// Package models
package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	TokenHash string    `json:"tokenHash" example:"$2a$10$6psfDpSdRdJRPnsZqWKH8eh0WLU9Bc6xKo7puJ27Z8s06te0AYqPa"`
	UserGUID  uuid.UUID `json:"userGuid" example:"a81bc81b-dead-4e5d-abff-90865d1e13b1"`
	CreatedAt time.Time `json:"createdAt" example:"2025-07-07T07:01:27.73104763Z"`
}

func RefreshTokenDBNew(hash string, userGUID uuid.UUID) RefreshToken {
	return RefreshToken{
		TokenHash: hash,
		UserGUID:  userGUID,
		CreatedAt: time.Now().UTC(),
	}
}

type InvalidAccessToken struct {
	GUID      uuid.UUID `json:"guid" example:"a81bc81b-dead-4e5d-abff-90865d1e13b1"`
	UserGUID  uuid.UUID `json:"userGuid" example:"a81bc81b-dead-4e5d-abff-90865d1e13b1"`
	CreatedAt time.Time `json:"createdAt" example:"2025-07-07T07:01:27.73104763Z"`
}

func InvalidAccessTokenDBNew(guid uuid.UUID, userGUID uuid.UUID) InvalidAccessToken {
	return InvalidAccessToken{
		GUID:      guid,
		UserGUID:  userGUID,
		CreatedAt: time.Now().UTC(),
	}
}

type LoginTokens struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NTE4ODM0NDIsInN1YiI6ImJhciJ9.1ZO0Znrv0CU1gYI52o7tTP1jHjzmpCi7ufyyStFcygeHWDCFZMya-12Uuswy_8saqxVs0mZx25hApJ3bpbPozA"`
	RefreshToken string `json:"refreshToken" example:"tmKpWaddK0DH5ldOEAQTUaWeH7hVabpH82cSLBzKw1ejXReJ2H+BHIRiLs7S3Nbm46c3sDko"`
}
