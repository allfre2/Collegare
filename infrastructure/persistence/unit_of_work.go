package persistence

import (
	"collegare/application/persistence"

	"gorm.io/gorm"
)

type BaseUnitOfWork struct {
	DB *gorm.DB
}

func (uow *BaseUnitOfWork) Begin() error {
	return uow.DB.Begin().Error
}

func (uow *BaseUnitOfWork) Commit() error {
	return uow.DB.Commit().Error
}

func (uow *BaseUnitOfWork) Rollback() error {
	return uow.DB.Rollback().Error
}

type CallinkUnitOfWork struct {
	BaseUnitOfWork
	linkRepo persistence.LinkRepository
}

func NewCallinkUnitOfWork(db *gorm.DB) persistence.CallinkUnitOfWork {
	return &CallinkUnitOfWork{
		BaseUnitOfWork: BaseUnitOfWork{
			DB: db,
		},
		linkRepo: NewLinkRepository(db),
	}
}

func (uow *CallinkUnitOfWork) LinkRepository() persistence.LinkRepository {
	return uow.linkRepo
}
