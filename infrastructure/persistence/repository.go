package persistence

import (
	"collegare/application/persistence"

	"gorm.io/gorm"
)

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
	return r.DB.Updates(&entity).Error
}

func (r *BaseRepository[T]) Delete(entity T) error {
	return r.DB.Delete(&entity, 0).Error
}
