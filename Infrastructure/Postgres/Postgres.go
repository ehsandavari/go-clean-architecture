package Postgres

import (
	"GolangCodeBase/Domain/Entities"
	"GolangCodeBase/Infrastructure"
	"context"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func init() {
	Infrastructure.Modules = append(Infrastructure.Modules, fx.Provide(NewPostgres))
}

type SPostgres struct {
	*gorm.DB
}

func NewPostgres(lc fx.Lifecycle, config SConfig) *SPostgres {
	sPostgres := new(SPostgres)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			if sPostgres.DB, err = gorm.Open(postgres.Open(config.URL), &gorm.Config{}); err != nil {
				return err
			}

			log.Println("postgres connection opened")
			return sPostgres.setup()
		},
		OnStop: func(ctx context.Context) error {
			log.Println("postgres connection closed")
			return sPostgres.close()
		},
	})
	return sPostgres
}

func (r *SPostgres) setup() error {
	return r.DB.AutoMigrate(
		new(Entities.OrderEntity),
	)
}

func (r *SPostgres) close() error {
	db, err := r.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (r *SPostgres) Transaction(fc func(tx *SPostgres) error) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		r.DB = tx
		return fc(r)
	})
}
