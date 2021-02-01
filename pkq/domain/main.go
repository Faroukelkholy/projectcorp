package domain

import (
	"projectcorp/config"
	http "projectcorp/pkq/adapter/http/client"
	"projectcorp/pkq/adapter/repository/postgres"
	usecase "projectcorp/pkq/domain/use_case"
	"sync"
)

var (
	instance *singleton
	once     sync.Once
)

type singleton struct {
	UseCases usecase.DomainUseCases
}

//nolint
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
		db := postgres.New(config.Parse())
		hc := &http.Client{}
		instance.UseCases = usecase.New(db, hc)
	})
	return instance
}
