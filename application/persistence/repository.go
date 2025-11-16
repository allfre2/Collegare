package persistence

import (
	"collegare/domain/models"

	"github.com/gofrs/uuid/v5"
)

type CallFlowRepository interface {
	FindAll() []models.CallFlow
	Find(uuid.UUID) models.CallFlow
}
