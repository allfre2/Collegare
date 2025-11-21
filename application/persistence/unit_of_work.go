package persistence

type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
}

type CallinkUnitOfWork interface {
	UnitOfWork
	LinkRepository() LinkRepository
}
