package persistence

import "gorm.io/gorm"

type PetRepository struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepository {
	return &PetRepository{
		db: db,
	}
}
