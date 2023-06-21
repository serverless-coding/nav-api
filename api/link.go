package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serverless-coding/frontend-nav/api/_pkg/db"
	"github.com/serverless-coding/frontend-nav/api/_pkg/dto"
	"github.com/serverless-coding/frontend-nav/api/_pkg/model"
	"github.com/serverless-coding/frontend-nav/api/_pkg/util"
)

func init() {
	util.Init()
}

func GetLinks() ([]*model.Link, error) {
	link := db.Use(db.GetDB()).ReadDB().Link
	links, err := link.WithContext(context.Background()).
		Where(link.Public.Is(true), link.Status.Eq(1)).Order(link.Rank).Find()
	if err != nil {
		return nil, err
	}

	return links, nil
}

type GetLinksRes struct {
	dto.BaseRespose
	Data []*model.Link `json:"data,omitempty"`
}

func Links(w http.ResponseWriter, r *http.Request) {
	ls, _ := GetLinks()
	res := GetLinksRes{
		BaseRespose: dto.Success(),
		Data:        ls,
	}
	jv, _ := json.Marshal(res)
	fmt.Fprint(w, string(jv))
}
