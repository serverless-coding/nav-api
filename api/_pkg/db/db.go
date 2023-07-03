package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	_db            *gorm.DB
	_once          sync.Once
	_slowThreshold = time.Millisecond * 100
)

type Stdout struct {
}

func (Stdout) Printf(s string, args ...interface{}) {
	fmt.Printf("%s\n[info] "+s, args...)
}

func Init(url string, _ ...*gorm.Option) {
	_once.Do(func() {
		var err error
		config := logger.Config{
			SlowThreshold: time.Duration(_slowThreshold),
			Colorful:      true,
			LogLevel:      logger.Info,
		}
		_db, err = gorm.Open(postgres.Open(url),
			&gorm.Config{
				Logger: logger.New(log.New(os.Stderr, "\r\n", log.LstdFlags), config),
			})
		if err != nil {
			panic(fmt.Errorf("open db %q fail: %w", url, err))
		}
	})
}

func GetDB(dbs ...*gorm.DB) *gorm.DB {
	if len(dbs) == 0 {
		return _db
	}
	return dbs[0]
}
