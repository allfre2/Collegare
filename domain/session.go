package domain

import (
	"github.com/gofrs/uuid/v5"
)

type SessionRequest struct {
	Context any
}

type Session struct {
	Id       uuid.UUID
	Outbound chan any
	IsActive bool
}

type AdHocSession struct {
	Id             uuid.UUID
	Caller         *Session
	Calle          *Session
	Offer          any
	CallerComplete bool
	CalleComplete  bool
}

func NewSession() *Session {
	id, _ := uuid.NewV7()
	return &Session{
		Id:       id,
		Outbound: make(chan any),
		IsActive: true,
	}
}

func (s *Session) Close() {
	close(s.Outbound)
	s.IsActive = false
}

func (s *AdHocSession) IsConnected() bool {
	return s.CallerComplete && s.CalleComplete
}

func (s *AdHocSession) GetPeer(session *Session) *Session {
	if s.Caller == session {
		return s.Calle
	} else if s.Calle == session {
		return s.Caller
	}

	return nil
}
