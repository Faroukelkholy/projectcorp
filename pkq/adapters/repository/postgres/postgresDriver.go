package postgres

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"projectcorp/config"
	"projectcorp/pkq/ports/output"
)

var cfg = config.Parse()

type PostgresDriver struct {
	pgConnection *pg.DB
	output.IProjectRepository
	output.IParticipantRepository
}


func NewPostgresDriver() output.IDatabasePort {
	postgresDriver := new(PostgresDriver)
	postgresDriver.connect()
	return postgresDriver
}


func (thisPS *PostgresDriver) connect() {
	thisPS.pgConnection = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%v", cfg.DBHost, cfg.DBPort),
		User:     cfg.DBUser,
		Password: cfg.DBPASS,
		Database: cfg.DBName,
	})
	thisPS.IProjectRepository = &ProjectRepository{Database: thisPS.pgConnection}
	thisPS.IParticipantRepository = &ParticipantRepository{Database: thisPS.pgConnection}
}