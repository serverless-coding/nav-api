package handler

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/serverless-coding/frontend-nav/api/util"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
)

func init() {
	yf, err := ioutil.ReadFile("../.env")
	if err != nil {
		panic(err)
	}
	var config interface{}
	yaml.Unmarshal(yf, &config)

	cfg, err := ini.Load("../.env")
	if err != nil {
		panic(err)
	}

	os.Setenv("DATABASE_URL", cfg.Section("").Key("DATABASE_URL").String())

	util.Init()
}

func TestLink(t *testing.T) {
	ls, err := GetLinks()
	if err != nil {
		t.Errorf("err:%+v", err)
	}

	fmt.Println(ls)
}
