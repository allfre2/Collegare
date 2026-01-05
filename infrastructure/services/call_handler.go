package services

import (
	"collegare/application/persistence"
	"collegare/application/services"
	"collegare/domain"

	"github.com/gofrs/uuid/v5"
)

// Custom parameters for webrtc
const (
	ICEServersParam         = "ICEServers"
	SessionDescriptionParam = "SessionDescription"
	IsOfferParam            = "IsOffer"
	IsCompleteParam         = "IsComplete"
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

		if !previouslyCreated {
			// Session was created now. Caller must wait for Calle to try to connect
			return nil, nil
		}

		if adHocSession.Calle == nil {
			adHocSession.Calle = session
			// Reply back with the offer
			cmd.Parameters[IsOfferParam] = true // (!) Remove this field ?
			cmd.Parameters[SessionDescriptionParam] = adHocSession.Offer

			// Respond with cached caller's offer
			session.Outbound <- cmd
			return nil, nil
		}

		// After this point just relay all messages until a completion message is received

		if ForwardIfNotConnected(cmd, adHocSession, session.Id) {
			(*s.sessionCache).RemoveAdHocRecord(adHocSession.Id) // Connection complete
		}

		return nil, nil
	}

	// Otherwise do nothing
	return nil, nil
}

func (s *WebRTCService) GetConfiguration() *domain.CallConfiguration {
	return &s.Configuration
}

func ForwardIfNotConnected(cmd *domain.CallCommand, adHocSession *domain.AdHocSession, sessionId uuid.UUID) bool {
	var destination *domain.Session

	isComplete := cmd.Parameters[IsCompleteParam]

	if sessionId == adHocSession.Calle.Id {
		adHocSession.CalleComplete = isComplete != nil && isComplete.(bool)
		destination = adHocSession.Caller
	}

	if sessionId == adHocSession.Caller.Id {
		adHocSession.CallerComplete = isComplete != nil && isComplete.(bool)
		destination = adHocSession.Calle
	}

	if destination == nil {
		return false
	}

	if !adHocSession.IsConnected() {
		destination.Outbound <- cmd
		return false
	}

	return true
}
