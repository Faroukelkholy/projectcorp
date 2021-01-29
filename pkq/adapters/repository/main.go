package repository

import (
	"projectcorp/pkq/adapters/repository/postgres"
	"projectcorp/pkq/ports/output"
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
