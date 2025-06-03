package domain

type Pagination[T any] struct {
	TotalItems int32 `json:"totalItems"`
	Page       int32 `json:"page"`
	Data       T     `json:"data"`
}
