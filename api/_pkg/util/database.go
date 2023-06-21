package util

import (
	"github.com/serverless-coding/frontend-nav/api/_pkg/config"
	"github.com/serverless-coding/frontend-nav/api/_pkg/db"
)

func Init() {
	conf := config.ParseConfig()
	db.Init(conf.DatabaseUrl)
}
