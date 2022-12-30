package models

import "golangCodeBase/domain/entities"

type DataModel[TE entities.IEntityConstraint] interface {
	ToEntity() TE
	FromEntity(entity TE) any
}
