package persistence

import (
	"collegare/domain"
	"collegare/domain/models"

	"collegare/application/persistence"

	"github.com/gofrs/uuid/v5"

	"gorm.io/gorm"
)

type MockCallFlowRepository struct{}

func NewMockCallFlowRepository() *MockCallFlowRepository {
	return &MockCallFlowRepository{}
}

func (r *MockCallFlowRepository) FindAll() []models.CallFlow {
	panic("Not Implemented")
}

func (r *MockCallFlowRepository) Find(id uuid.UUID) models.CallFlow {
	return models.CallFlow{
		Id:   id,
		Name: "Dummy flow",
		Script: models.Script{
			Start: 1,
			Nodes: map[int]models.Node{
				0: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "wilhelm.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
				1: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "c.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
				2: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "aeolian.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
				3: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "ionian.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
				4: {
					Id:   1,
					Loop: true,
					Name: "",
					Track: models.Track{
						Id:       uuid.Max,
						Filename: "mixolydian.opus",
					},
					Options: map[string]int{
						domain.One:   1,
						domain.Two:   2,
						domain.Three: 3,
						domain.Four:  4,
						domain.Zero:  0,
					},
				},
			},
		},
	}
}

type BaseRepository[T any] struct {
	persistence.Repository[T]
	DB *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		DB: db,
	}
}

func (r *BaseRepository[T]) Find(entity T) (T, error) {
	err := r.DB.First(&entity).Error
	return entity, err
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) Create(entity T) error {
	return r.DB.Create(&entity).Error
}

func (r *BaseRepository[T]) Update(entity T) error {
	return r.DB.Save(&entity).Error
}

func (r *BaseRepository[T]) Delete(entity T) error {
	return r.DB.Delete(&entity, 0).Error
}
