package util

import (
	"github.com/serverless-coding/nav-api/api/_pkg/config"
	"github.com/serverless-coding/nav-api/api/_pkg/db"
)

func Init() {
	conf := config.ParseConfig()
	if conf.DatabaseUrl == "" {
		return
	}
	db.Init(conf.DatabaseUrl)
}
