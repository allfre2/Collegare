package services

import "collegare/domain"

type CallHandler interface {
	Route(*domain.CallCommand, *domain.Session) (*domain.CallCommand, error)
	GetConfiguration() *domain.CallConfiguration
}
