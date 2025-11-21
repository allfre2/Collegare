package persistence

import (
	"collegare/domain/models"

	"github.com/gofrs/uuid/v5"
)

type CallFlowRepository interface {
	FindAll() []models.CallFlow
	Find(uuid.UUID) models.CallFlow
}

type Repository[T any] interface {
	Find(T) (T, error)
	FindAll() ([]T, error)
	Create(T) error
	Update(T) error
	Delete(T) error
}
