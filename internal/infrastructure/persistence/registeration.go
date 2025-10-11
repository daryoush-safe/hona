package persistence

import (
	"example/hona/bootstrap"
	"example/hona/internal/domain/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(dbConfig *bootstrap.Database) *Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Host,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Database{
		db: db,
	}
}

func (d *Database) AutoMigration() {
	d.db.AutoMigrate(
		&entities.User{},
		&entities.Pet{},
	)
}
