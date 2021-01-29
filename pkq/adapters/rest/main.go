package rest

import (
	"projectcorp/pkq/adapters/rest/gin"
	"projectcorp/pkq/ports/input"
	"sync"
)

var (
	instance *restAdapter
	once     sync.Once
)

type restAdapter struct {
	Adapter input.IRestAdapterPort
}

func NewRestAdapter() *restAdapter {
	once.Do(func() {
		instance = &restAdapter{}
		instance.Adapter = gin.NewGinServer()
	})
	return instance
}
