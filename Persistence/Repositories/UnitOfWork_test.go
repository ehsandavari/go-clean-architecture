package Repositories

import (
	"GolangCodeBase/Domain/Entities"
	DomainInterfaces "GolangCodeBase/Domain/Interfaces"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Suite struct {
	db         *SDatabaseContext
	mock       sqlmock.Sqlmock
	NameEntity Entities.OrderEntity
}

// todo: complate unit test
func TestUnitOfWork(t *testing.T) {
	s := &Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	dialector := mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       db,
	})
	s.db.Gorm, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	if s.db.Gorm == nil {
		t.Error("gorm db is null")
	}

	defer db.Close()

	var name = Entities.OrderEntity{
		Id:    1,
		Price: 100,
		Title: "order",
	}

	t.Run("Rollback on error", func(t *testing.T) {
		unitOfWork := NewUnitOfWork(s.db)
		err = unitOfWork.Do(func(work DomainInterfaces.IUnitOfWork) error {
			_ = work.OrderRepository().Add(name)
			require.NoError(t, err)

			name.Id = 2
			name.Price = 24
			name.Title = "test"
			_ = work.OrderRepository().Add(name)
			require.NoError(t, err)

			return nil
		})

		if assert.ErrorIs(t, err, nil) {
			find := unitOfWork.OrderRepository().Find()
			assert.EqualValues(t, find, nil)
		}
	})

}
