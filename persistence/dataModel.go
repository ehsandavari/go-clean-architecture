package persistence

import (
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
)

type IDataModel[TE entities.IEntityConstraint] interface {
	ToEntity() TE
	FromEntity(entity TE) any
}
