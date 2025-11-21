package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Link struct {
	Id         uuid.UUID `json:"id" gorm:"primary_key"`
	ExternalId string    `json:"external_id" gorm:"external_id"`
	Token      uuid.UUID `json:"token" gorm:"token"`
	Data       any       `json:"data" gorm:"data"`
	ExpiresAt  time.Time `json:"expires_at" gorm:"expires_at"`
	Auditable
}
