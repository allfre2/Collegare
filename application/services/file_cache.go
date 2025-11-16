package services

import (
	"collegare/domain/models"
)

type FileCache interface {
	GetTrack(models.Track) ([]byte, error)
}
