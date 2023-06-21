package db

import (
	"fmt"

	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	_db   *gorm.DB
	_once sync.Once
)

func Init(url string, _ ...*gorm.Option) {
	_once.Do(func() {
		var err error
		_db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("open db %q fail: %w", url, err))
		}
	})
}

func GetDB() *gorm.DB {
	return _db
}
