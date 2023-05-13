package models

import (
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (r *Base) ToEntity() entities.Base {
	return entities.Base{
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		IsDeleted: r.DeletedAt.Valid,
		DeletedAt: r.DeletedAt.Time,
	}
}

func (r *Base) FromEntity(entity entities.Base) {
	r.CreatedAt = entity.CreatedAt
	r.UpdatedAt = entity.UpdatedAt
	r.DeletedAt.Valid = entity.IsDeleted
	r.DeletedAt.Time = entity.DeletedAt
}
