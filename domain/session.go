package domain

import "github.com/gofrs/uuid/v5"

type SessionRequest struct {
	Context any
}

type Session struct {
	Id       uuid.UUID
	Inbound  chan *CallCommand
	Outbound chan *CallCommand
}

type AdHocSession struct {
	Id     uuid.UUID
	Caller *Session
	Calle  *Session
	Offer  any
}

func NewSession() *Session {
	id, _ := uuid.NewV7()
	return &Session{
		Id:       id,
		Inbound:  make(chan *CallCommand),
		Outbound: make(chan *CallCommand),
	}
}
