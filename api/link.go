package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/serverless-coding/frontend-nav/api/db"
	"github.com/serverless-coding/frontend-nav/api/dto"
	"github.com/serverless-coding/frontend-nav/api/model"
)

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
	Data []*model.Link
}

func Links(w http.ResponseWriter, r *http.Request) {
	ls, _ := GetLinks()
	res := GetLinksRes{
		Data: ls,
	}
	fmt.Fprint(w, res)
}
