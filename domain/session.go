package domain

import "github.com/gofrs/uuid/v5"

type SessionRequest struct {
	Context any
}

type Session struct {
	Id uuid.UUID
}

func NewSession() *Session {
	id, _ := uuid.NewV7()
	return &Session{
		Id: id,
	}
}
