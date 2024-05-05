package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStore struct {
	DB *gorm.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		DB: db,
	}, nil
}

func (ps *PostgresStore) Close() {
	db, _ := ps.DB.DB()
	db.Close()
}
