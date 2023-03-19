package persistence

import (
	"fmt"
	"github.com/ehsandavari/golang-clean-architecture/application/common"
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres/models"
	"github.com/google/uuid"
	"strings"
)

type sGenericRepository[TM models.IModel[TE], TE entities.IEntityConstraint] struct {
	dataBaseContext *SDatabaseContext
}

func newGenericRepository[TM models.IModel[TE], TE entities.IEntityConstraint](dataBaseContext *SDatabaseContext) interfaces.IGenericRepository[TE] {
	return sGenericRepository[TM, TE]{
		dataBaseContext: dataBaseContext,
	}
}

func (r sGenericRepository[TM, TE]) Paginate(listQuery common.PaginateQuery) (*common.PaginateResult[TE], error) {
	var model TM
	var totalRows int64
	r.dataBaseContext.Postgres.Model(model).Count(&totalRows)
	query := r.dataBaseContext.Postgres.Model(model).Offset(listQuery.GetOffset()).Limit(listQuery.GetLimit()).Order(listQuery.GetOrderBy())
	if listQuery.Filters != nil {
		for _, filter := range listQuery.Filters {
			column := filter.Key
			action := filter.Comparison
			value := filter.Value

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				query = query.Where(whereQuery, value)
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				query = query.Where(whereQuery, "%"+value+"%")
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(value, ",")
				query = query.Where(whereQuery, queryArray)
			}
		}
	}

	var entitiesObjects []TE
	if err := query.Find(&entitiesObjects).Error; err != nil {
		return nil, err
	}

	return common.NewPaginateResult[TE](entitiesObjects, listQuery.GetPage(), listQuery.GetPerPage(), totalRows), nil
}

func (r sGenericRepository[TM, TE]) All() []TE {
	var model TM
	var entitiesObjects []TE
	r.dataBaseContext.Postgres.DB.Model(model).Find(&entitiesObjects)
	return entitiesObjects
}

func (r sGenericRepository[TM, TE]) First() TE {
	var model TM
	r.dataBaseContext.Postgres.DB.First(&model)
	return model.ToEntity()
}

func (r sGenericRepository[TM, TE]) Last() TE {
	var model TM
	r.dataBaseContext.Postgres.DB.Last(&model)
	return model.ToEntity()
}

func (r sGenericRepository[TM, TE]) Add(entity TE) int64 {
	var model TM
	model = model.FromEntity(entity).(TM)
	return r.dataBaseContext.Postgres.DB.Create(&model).RowsAffected
}

func (r sGenericRepository[TM, TE]) Update(id uuid.UUID, entity TE) int64 {
	var model TM
	model = model.FromEntity(entity).(TM)
	return r.dataBaseContext.Postgres.DB.Where("id", id).Updates(&model).RowsAffected
}

func (r sGenericRepository[TM, TE]) Delete(id uuid.UUID) int64 {
	var model TM
	return r.dataBaseContext.Postgres.DB.Delete(&model, id).RowsAffected
}
