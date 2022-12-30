package persistence

import (
	"golangCodeBase/domain/entities"
)

type (
	IDataModel[TE entities.IEntityConstraint] interface {
		ToEntity() TE
		FromEntity(entity TE) any
	}
)
