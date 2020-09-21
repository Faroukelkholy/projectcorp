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
		//clientRestAdapter := clientRest.ClientRest{}
		//instance.DomainUseCases = useCases.NewDomainUseCases(databaseAdapter.Adapter,clientRestAdapter)
	})
	return instance
}


