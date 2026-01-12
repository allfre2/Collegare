package services

import (
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
	sessionCache  *services.SessionCache
}

func NewWebRTCService(ICEServers []string, sessionCache *services.SessionCache) *WebRTCService {
	return &WebRTCService{
		ICEServers: ICEServers,
		Configuration: domain.CallConfiguration{
			Protocol: domain.WebRTC,
			Parameters: map[string]any{
				ICEServersParam: ICEServers,
			},
		},
		sessionCache: sessionCache,
	}
}

func (s *WebRTCService) Route(cmd *domain.CallCommand, session *domain.Session) error {
	// Perform signaling between callers
	adHocSession, previouslyCreated, err := (*s.sessionCache).ReadInsertAdHocSession(cmd.AdHocId, &domain.AdHocSession{
		Id:     cmd.AdHocId,
		Caller: session,
		Offer:  cmd.Parameters[SessionDescriptionParam],
	})
	if err != nil {
		return err
	}

	if !previouslyCreated {
		// Session was created now. Caller must wait for Calle to try to connect
		return nil
	}

	if adHocSession.Calle == nil {
		adHocSession.Calle = session
		// Reply back with the offer
		cmd.Parameters[IsOfferParam] = true // (!) Remove this field ?
		cmd.Parameters[SessionDescriptionParam] = adHocSession.Offer

		// Respond with cached caller's offer
		session.Outbound <- cmd
		return nil
	}

	// After this point just relay all messages until a completion message is received
	relayUntilConnected(cmd, adHocSession, session.Id)

	return nil
}

func (s *WebRTCService) EndCall(cmd *domain.HangUpCommand, session *domain.Session) error {
	if adHocSession := (*s.sessionCache).GetAdHocSession(cmd.AdHocId); adHocSession != nil {
		if peer := adHocSession.GetPeer(session); peer != nil {
			peer.Outbound <- cmd
		}

		(*s.sessionCache).RemoveAdHocRecord(adHocSession.Id)
	}

	return nil
}

func (s *WebRTCService) GetConfiguration() *domain.CallConfiguration {
	return &s.Configuration
}

func relayUntilConnected(cmd *domain.CallCommand, adHocSession *domain.AdHocSession, sessionId uuid.UUID) bool {
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
