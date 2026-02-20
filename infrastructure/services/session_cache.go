package services

import (
	"collegare/domain"
	"sync"

	"github.com/gofrs/uuid/v5"
)

type DefaultSessionCache struct {
	cache      map[string]*domain.Session
	adHocCache map[string]*domain.AdHocSession
	mutex      sync.RWMutex
}

func NewSessionCache() *DefaultSessionCache {
	return &DefaultSessionCache{
		cache:      make(map[string]*domain.Session),
		adHocCache: make(map[string]*domain.AdHocSession),
	}
}

func (c *DefaultSessionCache) GetSession(uuid.UUID) *domain.Session {
	panic("Not Implemented")
}

func (c *DefaultSessionCache) GetAdHocSession(id uuid.UUID) *domain.AdHocSession {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.getAdHocSession(id)
}

func (c *DefaultSessionCache) CreateAdHocSession(session *domain.AdHocSession) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.writeAdHocSession(session)
}

func (c *DefaultSessionCache) ReadInsertAdHocSession(id uuid.UUID, new *domain.AdHocSession) (session *domain.AdHocSession, isPresent bool, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if session := c.getAdHocSession(id); session != nil {
		return session, true, nil
	}

	if err := c.writeAdHocSession(new); err != nil {
		return nil, false, err
	}

	return new, false, nil
}

func (c *DefaultSessionCache) RemoveAdHocRecord(id uuid.UUID) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.adHocCache, id.String())

	return nil
}

func (c *DefaultSessionCache) getAdHocSession(id uuid.UUID) *domain.AdHocSession {
	if session, ok := c.adHocCache[id.String()]; ok {
		return session
	}

	return nil
}

func (c *DefaultSessionCache) writeAdHocSession(session *domain.AdHocSession) error {
	c.adHocCache[session.Id.String()] = session
	return nil
}
