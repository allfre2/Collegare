package persistence

import (
	"collegare/application/persistence"
	"collegare/domain/models"

	"gorm.io/gorm"
)

type LinkRepository struct {
	BaseRepository[models.Link]
}

func NewLinkRepository(db *gorm.DB) persistence.LinkRepository {
	return &LinkRepository{
		BaseRepository: BaseRepository[models.Link]{
			DB: db,
		},
	}
}
