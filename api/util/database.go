package util

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{}))
	if err != nil {
		return nil, err
	}

	return db, nil
}
