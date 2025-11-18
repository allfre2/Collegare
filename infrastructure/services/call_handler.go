package services

import (
	"collegare/application/persistence"
	"collegare/application/services"
	"collegare/domain"
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
	sessionCache  *services.SessionCache
}

func NewWebRTCService(ICEServers []string, callFlowRepo persistence.CallFlowRepository, fileCache services.FileCache, sessionCache *services.SessionCache) *WebRTCService {
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
	if cmd.Mode == domain.AdHoc {
		// Perform signaling between callers
		adHocSession, previouslyCreated, err := (*s.sessionCache).ReadInsertAdHocSession(cmd.AdHocId, &domain.AdHocSession{
			Id:     cmd.AdHocId,
			Caller: session,
			Offer:  cmd.Parameters[SessionDescriptionParam],
		})
		if err != nil {
			return nil, err
		}
		// fmt.Println("AdHoc Session: ", adHocSession, previouslyCreated)
		if !previouslyCreated {
			// Session was created now. Caller must wait for Calle to try to connect
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

			(*s.sessionCache).RemoveAdHocRecord(adHocSession.Id)

			return nil, nil
		}
	}

	// Otherwise do nothing
	return nil, nil
}

func (s *WebRTCService) GetConfiguration() *domain.CallConfiguration {
	return &s.Configuration
}
