package services

import (
	"collegare/application/persistence"
	"collegare/application/services"
	"collegare/domain"
)

type WebRTCService struct {
	STUNServers   []string
	Configuration domain.CallConfiguration
	callFlowRepo  persistence.CallFlowRepository
	cache         services.FileCache
}

func NewWebRTCService(stunServers []string, callFlowRepo persistence.CallFlowRepository, cache services.FileCache) *WebRTCService {
	return &WebRTCService{
		STUNServers: stunServers,
		Configuration: domain.CallConfiguration{
			Protocol: domain.WebRTC,
			Parameters: map[string]any{
				"stunServers": stunServers,
			},
		},
		callFlowRepo: callFlowRepo,
		cache:        cache,
	}
}

func (wbrtcs *WebRTCService) GetConfiguration() *domain.CallConfiguration {
	return &wbrtcs.Configuration
}
