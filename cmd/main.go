package main

import (
	"log"
	"projectcorp/config"
	"projectcorp/pkq/adapters/rest"
)

func main(){
	restAdapter := rest.NewRestAdapter()
	log.Println("portttt:", config.GetEnvConfig().HTTP_PORT)
	err := restAdapter.Adapter.Start(config.GetEnvConfig().HTTP_PORT)

	if err != nil {
		log.Fatal(err)
	}
}