package interfaces

import (
	"context"
	"github.com/ehsandavari/go-clean-architecture/application/common"
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
	"github.com/google/uuid"
)

//go:generate mockgen -destination=../mocks/mockPersistence.go -package=mocks github.com/ehsandavari/go-clean-architecture/application/common/interfaces IUnitOfWork

type (
	IGenericRepository[TE entities.IEntityConstraint] interface {
		Paginate(ctx context.Context, listQuery *common.PaginateQuery) (*common.PaginateResult[TE], error)
		All(ctx context.Context) ([]TE, error)
		FirstById(ctx context.Context, id uuid.UUID) (*TE, error)
		First(ctx context.Context) (*TE, error)
		Last(ctx context.Context) (*TE, error)
		Create(ctx context.Context, entity *TE) (int64, error)
		Update(ctx context.Context, id uuid.UUID, entity *TE) (int64, error)
		Delete(ctx context.Context, id uuid.UUID) (int64, error)
	}
	ISongRepository interface {
		IGenericRepository[entities.Song]
	}
	IUnitOfWork interface {
		SongRepository() ISongRepository
		Do(func(IUnitOfWork) error) error
	}
)
