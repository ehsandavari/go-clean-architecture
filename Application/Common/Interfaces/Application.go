package Interfaces

import (
	"context"
)

type (
	IBaseCommand[T any] interface {
		Handle(ctx context.Context, command T) error
	}
)
