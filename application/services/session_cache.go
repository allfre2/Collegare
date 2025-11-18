package services

import (
	"collegare/domain"

	"github.com/gofrs/uuid/v5"
)

type SessionCache interface {
	GetSession(uuid.UUID) *domain.Session
	GetAdHocSession(uuid.UUID) *domain.AdHocSession
	CreateAdHocSession(*domain.AdHocSession) error
}
