package services

import "collegare/domain"

type CommandHandler interface {
	Parse(string) (any, error)
	Handle(string, *domain.Session) error
	GetCallConfiguration() (*domain.CallConfiguration, error)
}
