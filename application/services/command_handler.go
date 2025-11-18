package services

import "collegare/domain"

type CommandHandler interface {
	Parse(string) (any, error)
	Handle(string, *domain.Session) (string, error)
	GetCallConfiguration() (*domain.CallConfiguration, error)
}
