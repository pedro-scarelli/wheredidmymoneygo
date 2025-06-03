package dto

import (
	"net/http"
	"strconv"
)

type PaginationRequestParams struct {
	Page         int `form:"page"`
	ItemsPerPage int `form:"itemsPerPage"`
}

func FromValuePaginationRequestParams(request *http.Request) (*PaginationRequestParams, error) {
	query := request.URL.Query()

	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}

	itemsPerPage, _ := strconv.Atoi(query.Get("itemsPerPage"))
	if itemsPerPage < 1 {
		itemsPerPage = 10
	}

	return &PaginationRequestParams{
		Page:         page,
		ItemsPerPage: itemsPerPage,
	}, nil

}
