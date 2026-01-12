package services

import "collegare/domain"

type CallHandler interface {
	Route(*domain.CallCommand, *domain.Session) error
	EndCall(*domain.HangUpCommand, *domain.Session) error
	GetConfiguration() *domain.CallConfiguration
}
