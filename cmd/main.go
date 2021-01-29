package main

import (
	"log"
	"projectcorp/config"
	"projectcorp/pkq/adapters/rest"
)

func main(){
	restAdapter := rest.NewRestAdapter()
	err := restAdapter.Adapter.Start(config.Parse().HTTPPort)

	if err != nil {
		log.Fatal(err)
	}
}