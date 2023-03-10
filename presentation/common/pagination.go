package common

import (
	"math"
)

const (
	defaultSize = 10
	defaultPage = 1
)

type ListResult[T any] struct {
	Size       int   `json:"size,omitempty" bson:"size"`
	Page       int   `json:"page,omitempty" bson:"page"`
	TotalItems int64 `json:"totalItems,omitempty" bson:"totalItems"`
	TotalPage  int   `json:"totalPage,omitempty" bson:"totalPage"`
	Items      []T   `json:"items,omitempty" bson:"items"`
}

func NewListResult[T any](items []T, size int, page int, totalItems int64) *ListResult[T] {
	listResult := &ListResult[T]{Items: items, Size: size, Page: page, TotalItems: totalItems}

	listResult.TotalPage = getTotalPages(totalItems, size)

	return listResult
}

func getTotalPages(totalCount int64, size int) int {
	d := float64(totalCount) / float64(size)
	return int(math.Ceil(d))
}

type FilterModel struct {
	Field      string `query:"field" json:"field"`
	Value      string `query:"value" json:"value"`
	Comparison string `query:"comparison" json:"comparison"`
}

type ListQuery struct {
	Size    int           `query:"size" json:"size,omitempty"`
	Page    int           `query:"page" json:"page,omitempty"`
	OrderBy string        `query:"orderBy" json:"orderBy,omitempty"`
	Filters []FilterModel `query:"filters" json:"filters,omitempty"`
}

func NewListQuery(size int, page int) *ListQuery {
	return &ListQuery{Size: size, Page: page}
}

func (q *ListQuery) SetSize(size int) {
	if size == 0 {
		q.Size = defaultSize
	}
	q.Size = size
}

func (q *ListQuery) SetPage(page int) {
	if page == 0 {
		q.Page = defaultPage
	}
	q.Page = page
}

func (q *ListQuery) SetOrderBy(orderByQuery string) {
	q.OrderBy = orderByQuery
}

func (q ListQuery) GetOffset() int {
	return (q.Page - 1) * q.Size
}

func (q ListQuery) GetLimit() int {
	return q.Size
}

func (q ListQuery) GetOrderBy() string {
	return q.OrderBy
}

func (q ListQuery) GetPage() int {
	return q.Page
}

func (q ListQuery) GetSize() int {
	return q.Size
}
