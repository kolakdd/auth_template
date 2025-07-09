// Package models
package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type WebHookDto struct {
	UserID uuid.UUID `json:"userID" example:" a81bc81b-dead-4e5d-abff-90865d1e13b1"`
	OldIP  string    `json:"OldIP" example:"127.0.0.1"`
	NewIP  string    `json:"NewIP" example:"127.0.0.2"`
	Time   time.Time `json:"time" example:"2025-07-07T07:01:27.73104763Z"`
}

func (whd WebHookDto) String() string {
	return fmt.Sprintf("userId=%s, oldIP=%s,newIp=%s, time=%s", whd.UserID, whd.OldIP, whd.NewIP, whd.Time)
}
