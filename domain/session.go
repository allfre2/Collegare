package domain

import "github.com/gofrs/uuid/v5"

type SessionRequest struct {
	Context any
}

type Session struct {
	Id       uuid.UUID
	Outbound chan any
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
	id, _ := uuid.NewV7() // handle errors
	return &Session{
		Id:       id,
		Outbound: make(chan any),
	}
}

func (s *Session) Close() {
	close(s.Outbound)
}

func (s *AdHocSession) IsConnected() bool {
	return s.CallerComplete && s.CalleComplete
}
