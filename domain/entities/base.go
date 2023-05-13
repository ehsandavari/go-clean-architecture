package entities

import (
	"time"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
	DeletedAt time.Time
}

func NewBase(createdAt time.Time, updatedAt time.Time) Base {
	return Base{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (r *Base) SetDeleted(isDeleted bool) {
	r.IsDeleted = isDeleted
	if isDeleted {
		r.DeletedAt = time.Now()
	}
}
