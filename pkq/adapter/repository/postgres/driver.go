package postgres

import (
	"fmt"
	"log"
	"projectcorp/config"
	"projectcorp/pkq/port/output"

	"github.com/go-pg/pg/v9"
)


type Driver struct {
	DB *pg.DB
	output.IProjectRepository
	output.IParticipantRepository
}

func New(cfg *config.Config) output.IRepositoryPort {
	d := new(Driver)
	d.connect(cfg)
	d.IProjectRepository = &ProjectRepository{DB: d.DB}
	d.IParticipantRepository = &ParticipantRepository{DB: d.DB}
	return d
}

func (d *Driver) connect(cfg *config.Config) {
	log.Print("conf.DB_HOST :", cfg.DBHost)
	d.DB = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%v", cfg.DBHost, cfg.DBPort),
		User:     cfg.DBUser,
		Password: cfg.DBPASS,
		Database: cfg.DBName,
	})
}
