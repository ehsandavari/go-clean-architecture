package persistence

import (
	"golangCodeBase/domain/entities"
	"golangCodeBase/infrastructure/postgres/models"
)

type (
	IModelConstraint interface {
		models.OrderModel
	}
	IDataModel[TE entities.IEntityConstraint] interface {
		ToEntity() TE
		FromEntity(entity TE) any
	}
)
