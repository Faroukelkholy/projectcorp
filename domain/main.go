package domain

import (
	"projectcorp/domain/useCases"
	"sync"
)

var (
	instance *DomainUseCasesSingleton
	once     sync.Once
)

type DomainUseCasesSingleton struct {
	DomainUseCases useCases.DomainUseCases
}

func NewDomainUseCases() *DomainUseCasesSingleton {
	once.Do(func() {
		instance = &DomainUseCasesSingleton{}
		//databaseAdapter := repository.NewDatabaseAdapter()
		//restAdapter := restClient{}
		//instance.DomainUseCases = useCases.NewDomainUseCases(databaseAdapter.Adapter,restAdapter)
	})
	return instance
}


