package services

import (
	"collegare/domain"
)

type SessionService interface {
	Connect(*domain.SessionRequest) (*domain.Session, error)
	Disconnect(*domain.Session) error
	GetSessionConfiguration() (*domain.ConfigurationMessage, error)
}
