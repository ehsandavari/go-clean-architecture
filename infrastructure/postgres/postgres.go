package postgres

import (
	"github.com/ehsandavari/go-clean-architecture/infrastructure/postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type SPostgres struct {
	*gorm.DB
}

func NewPostgres(config *SConfig) *SPostgres {
	sPostgres := new(SPostgres)
	var err error

	if sPostgres.DB, err = gorm.Open(postgres.Open("host=" + config.Host + " user=" + config.User + " password=" + config.Password + " dbname=" + config.DatabaseName + " port=" + config.Port + " sslmode=" + config.SslMode + " TimeZone=" + config.TimeZone + "")); err != nil {
		log.Fatalln("error in connect to postgres ", err)
	}

	log.Println("postgres connection opened")

	if err = sPostgres.setup(); err != nil {
		log.Fatalln("error in setup postgres ", err)
	}

	return sPostgres
}

func (r *SPostgres) setup() error {
	return r.DB.AutoMigrate(
		new(models.Song),
		new(models.SongTariff),
	)
}

func (r *SPostgres) Close() error {
	db, err := r.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (r *SPostgres) Transaction(fc func(*SPostgres) error) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		r.DB = tx
		return fc(r)
	})
}
