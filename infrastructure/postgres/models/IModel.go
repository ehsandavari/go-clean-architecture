package models

import (
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
)

type IModel[TE entities.IEntityConstraint] interface {
	ToEntity() *TE
	FromEntity(entity *TE) any
}
