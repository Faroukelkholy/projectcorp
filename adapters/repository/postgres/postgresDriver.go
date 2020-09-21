package postgres

import (
	"fmt"
	"projectcorp/ports/output"
	"projectcorp/utils"
	"github.com/go-pg/pg/v9"
)

var config = utils.GetEnvConfig()

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
		Addr:     fmt.Sprintf("%s:%v", config.DB_HOST, config.DB_PORT),
		User:     config.DB_USERNAME,
		Password: config.DB_PASSWORD,
		Database: config.DB_NAME,
	})
	thisPS.IProjectRepository = &ProjectRepository{Database: thisPS.pgConnection}
	thisPS.IParticipantRepository = &ParticipantRepository{Database: thisPS.pgConnection}
}