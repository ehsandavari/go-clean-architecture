package GetAllOrderByFilter

import "github.com/ehsandavari/golang-clean-architecture/application/common"

type SGetAllOrderByFilterQuery struct {
	common.PaginateQuery
}

func NewSGetAllOrderByFilterQuery(paginateQuery common.PaginateQuery) SGetAllOrderByFilterQuery {
	return SGetAllOrderByFilterQuery{
		PaginateQuery: paginateQuery,
	}
}
