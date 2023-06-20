package util

import (
	"github.com/serverless-coding/frontend-nav/api/config"
	"github.com/serverless-coding/frontend-nav/api/db"
)

func Init() {
	conf := config.ParseConfig()
	db.Init(conf.DatabaseUrl)
}
