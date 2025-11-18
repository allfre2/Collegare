package services

import (
	"collegare/domain"

	"github.com/gofrs/uuid/v5"
)

type MockSessionCache struct {
	cache      map[string]*domain.Session
	adHocCache map[string]*domain.AdHocSession
}

func NewMockSessionCache() *MockSessionCache {
	return &MockSessionCache{
		cache:      make(map[string]*domain.Session),
		adHocCache: make(map[string]*domain.AdHocSession),
	}
}

func (c *MockSessionCache) GetSession(uuid.UUID) *domain.Session {
	panic("Not Implemented")
}

func (c *MockSessionCache) GetAdHocSession(id uuid.UUID) *domain.AdHocSession {
	if session, ok := c.adHocCache[id.String()]; ok {
		return session
	}

	return nil
}

func (c *MockSessionCache) CreateAdHocSession(session *domain.AdHocSession) error {
	c.adHocCache[session.Id.String()] = session
	return nil
}
