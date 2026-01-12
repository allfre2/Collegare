package persistence

type Repository[T any] interface {
	Find(T) (T, error)
	FindAll() ([]T, error)
	Create(T) error
	Update(T) error
	Delete(T) error
}
