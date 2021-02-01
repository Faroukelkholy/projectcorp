package main

import (
	"log"
	"projectcorp/config"
	"projectcorp/pkq/adapter/http/server"
	"projectcorp/pkq/domain"
)

func main() {
	singleton := domain.GetInstance()
	s := server.New(singleton.UseCases)
	err := s.Start(config.Parse().HTTPPort)

	if err != nil {
		log.Fatal(err)
	}
}
