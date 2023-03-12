package common

import (
	"math"
)

const (
	defaultPage    = 1
	defaultPerPage = 10
)

type PaginateResult[T any] struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	TotalPage  int   `json:"total_page"`
	TotalItems int64 `json:"total_items"`
	Items      []T   `json:"items"`
} //@name PaginateResult

func NewPaginateResult[T any](items []T, page int, perPage int, totalItems int64) *PaginateResult[T] {
	paginateResult := &PaginateResult[T]{Items: items, Page: page, PerPage: perPage, TotalItems: totalItems}

	paginateResult.TotalPage = getTotalPages(totalItems, perPage)

	return paginateResult
}

func getTotalPages(totalCount int64, perPage int) int {
	d := float64(totalCount) / float64(perPage)
	return int(math.Ceil(d))
}

type ComparisonType string //@name ComparisonType

const (
	In       ComparisonType = "in"
	Equals   ComparisonType = "equals"
	Contains ComparisonType = "contains"
)

func (r ComparisonType) String() string {
	return string(r)
}

type FilterQuery struct {
	Key        string         `json:"key"`
	Comparison ComparisonType `json:"comparison" enums:"in,equals,contains"`
	Value      string         `json:"value"`
} //@name FilterQuery

type PaginateQuery struct {
	Page    int           `json:"page" example:"1" default:"1"`
	PerPage int           `json:"per_page" example:"10" default:"10"`
	OrderBy string        `json:"order_by"`
	Filters []FilterQuery `json:"filters"`
} //@name PaginateQuery

func NewPaginateQuery(perPage int, page int) PaginateQuery {
	return PaginateQuery{PerPage: perPage, Page: page}
}

func (q *PaginateQuery) SetPerPage(perPage int) {
	if perPage == 0 {
		q.PerPage = defaultPerPage
	}
	q.PerPage = perPage
}

func (q *PaginateQuery) SetPage(page int) {
	if page == 0 {
		q.Page = defaultPage
	}
	q.Page = page
}

func (q *PaginateQuery) SetOrderBy(orderByQuery string) {
	q.OrderBy = orderByQuery
}

func (q *PaginateQuery) GetOffset() int {
	return (q.Page - 1) * q.PerPage
}

func (q *PaginateQuery) GetLimit() int {
	return q.PerPage
}

func (q *PaginateQuery) GetOrderBy() string {
	return q.OrderBy
}

func (q *PaginateQuery) GetPage() int {
	return q.Page
}

func (q *PaginateQuery) GetPerPage() int {
	return q.PerPage
}
