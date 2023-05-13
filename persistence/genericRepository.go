package persistence

import (
	"context"
	"fmt"
	"github.com/ehsandavari/go-clean-architecture/application/common"
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
	"github.com/ehsandavari/go-clean-architecture/infrastructure/postgres/models"
	"github.com/google/uuid"
	"strings"
)

type sGenericRepository[TM models.IModel[TE], TE entities.IEntityConstraint] struct {
	*sDatabaseContext
}

func newGenericRepository[TM models.IModel[TE], TE entities.IEntityConstraint](dataBaseContext *sDatabaseContext) sGenericRepository[TM, TE] {
	return sGenericRepository[TM, TE]{
		sDatabaseContext: dataBaseContext,
	}
}

func (r sGenericRepository[TM, TE]) Paginate(ctx context.Context, listQuery *common.PaginateQuery) (*common.PaginateResult[TE], error) {
	var model TM
	var totalRows int64
	r.Postgres.WithContext(ctx).Model(model).Count(&totalRows)
	query := r.Postgres.WithContext(ctx).Model(model).Offset(listQuery.GetOffset()).Limit(listQuery.GetLimit()).Order(listQuery.GetOrderBy())
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

func (r sGenericRepository[TM, TE]) All(ctx context.Context) ([]TE, error) {
	var model TM
	var entitiesObjects []TE
	result := r.Postgres.WithContext(ctx).Model(model).Find(&entitiesObjects)
	if result.Error != nil {
		return nil, result.Error
	}
	return entitiesObjects, nil
}

func (r sGenericRepository[TM, TE]) FirstById(ctx context.Context, id uuid.UUID) (*TE, error) {
	var model TM
	result := r.Postgres.WithContext(ctx).First(&model, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return model.ToEntity(), nil
}

func (r sGenericRepository[TM, TE]) First(ctx context.Context) (*TE, error) {
	var model TM
	result := r.Postgres.WithContext(ctx).First(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model.ToEntity(), nil
}

func (r sGenericRepository[TM, TE]) Last(ctx context.Context) (*TE, error) {
	var model TM
	result := r.Postgres.WithContext(ctx).Last(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model.ToEntity(), nil
}

func (r sGenericRepository[TM, TE]) Create(ctx context.Context, entity *TE) (int64, error) {
	var model TM
	model = model.FromEntity(entity).(TM)
	result := r.Postgres.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (r sGenericRepository[TM, TE]) Update(ctx context.Context, id uuid.UUID, entity *TE) (int64, error) {
	var model TM
	model = model.FromEntity(entity).(TM)
	result := r.Postgres.WithContext(ctx).Where("id", id).Updates(&model)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (r sGenericRepository[TM, TE]) Delete(ctx context.Context, id uuid.UUID) (int64, error) {
	var model TM
	result := r.Postgres.WithContext(ctx).Delete(&model, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
