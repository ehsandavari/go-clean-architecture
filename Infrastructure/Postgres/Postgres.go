package Postgres

import (
	DomainEntities "GolangCodeBase/Domain/Entities"
	"context"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type SPostgres struct {
	DB *gorm.DB
}

func NewPostgres(lc fx.Lifecycle, config *SConfig) *SPostgres {
	sPostgres := new(SPostgres)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			if sPostgres.DB, err = gorm.Open(postgres.Open(config.URL), &gorm.Config{}); err != nil {
				return err
			}

			log.Println("postgres database connection successfully")
			return sPostgres.setup()
		},
		OnStop: func(ctx context.Context) error {
			log.Println("postgres database connection closed")
			return sPostgres.close()
		},
	})
	return sPostgres
}

func (r *SPostgres) setup() error {
	return r.DB.AutoMigrate(
		new(DomainEntities.OrderEntity),
	)
}

func (r *SPostgres) close() error {
	db, err := r.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
