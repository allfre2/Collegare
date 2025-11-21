package persistence

import (
	"collegare/domain/models"
)

type LinkRepository interface {
	Repository[models.Link]
}
