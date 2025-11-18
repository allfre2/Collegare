package services

import (
	"collegare/application/persistence"
	"collegare/application/services"
	"collegare/domain"
	"fmt"
)

// Custom parameters for webrtc
const (
	ICEServersParam         = "ICEServers"
	SessionDescriptionParam = "SessionDescription"
	IsOfferParam            = "IsOffer"
)

type WebRTCService struct {
	ICEServers    []string
	Configuration domain.CallConfiguration
	callFlowRepo  persistence.CallFlowRepository
	fileCache     services.FileCache
	sessionCache  services.SessionCache
}

func NewWebRTCService(ICEServers []string, callFlowRepo persistence.CallFlowRepository, fileCache services.FileCache, sessionCache services.SessionCache) *WebRTCService {
	return &WebRTCService{
		ICEServers: ICEServers,
		Configuration: domain.CallConfiguration{
			Protocol: domain.WebRTC,
			Parameters: map[string]any{
				ICEServersParam: ICEServers,
			},
		},
		callFlowRepo: callFlowRepo,
		fileCache:    fileCache,
		sessionCache: sessionCache,
	}
}

func (s *WebRTCService) Route(cmd *domain.CallCommand, session *domain.Session) (*domain.CallCommand, error) {
	// Signaling code ...
	if cmd.Mode == domain.AdHoc {
		// Perform signaling between callers
		adHocSession := s.sessionCache.GetAdHocSession(cmd.AdHocId)
		if adHocSession == nil {
			s.sessionCache.CreateAdHocSession(&domain.AdHocSession{
				Id:     cmd.AdHocId,
				Caller: session,
				Offer:  cmd.Parameters[SessionDescriptionParam],
			})

			return nil, nil
		} else if adHocSession.Calle == nil {
			adHocSession.Calle = session
			// Reply back with the offer
			cmd.Parameters[IsOfferParam] = true
			cmd.Parameters[SessionDescriptionParam] = adHocSession.Offer

			// Respond with cached caller's offer
			return cmd, nil
		} else if adHocSession.Calle.Id == session.Id {
			// Now It's an answer from the Calle
			adHocSession.Caller.Outbound <- cmd
			// Answer has been seent
			adHocSession.Answer = cmd.Parameters[SessionDescriptionParam]
			fmt.Println("AdHocSession: ", adHocSession)

			return nil, nil
		}
	}

	// Otherwise do nothing
	return nil, nil
}

func (s *WebRTCService) GetConfiguration() *domain.CallConfiguration {
	return &s.Configuration
}
