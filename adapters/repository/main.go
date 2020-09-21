package repository

import (
	"projectcorp/adapters/repository/postgres"
	"projectcorp/ports/output"
	"sync"
)

var (
	instance *databaseAdapter
	once     sync.Once
)

type databaseAdapter struct {
	Adapter output.IDatabasePort
}

func NewDatabaseAdapter() *databaseAdapter {
	once.Do(func() {
		instance = &databaseAdapter{}
		instance.Adapter = postgres.NewPostgresDriver()
	})
	return instance
}
