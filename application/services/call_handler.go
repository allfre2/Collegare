package services

import "collegare/domain"

type CallHandler interface {
	GetConfiguration() *domain.CallConfiguration
}
