package persistence

import "gorm.io/gorm"

type RepositoryFactory struct {
	db *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{
		db: db,
	}
}

func (f *RepositoryFactory) UserRepository() *UserRepository {
	return NewUserRepository(f.db)
}

func (f *RepositoryFactory) PetRepository() *PetRepository {
	return NewPetRepository(f.db)
}
