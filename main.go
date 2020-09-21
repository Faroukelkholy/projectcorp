package main

import (
	"log"
	"projectcorp/adapters/rest"
	"projectcorp/utils"
)

func main(){

	log.Println("GETEMPLOYEES_URL:", utils.GetEnvConfig().GETEMPLOYEES_URL)
	restAdapter := rest.NewRestAdapter()
	log.Println("portttt:", utils.GetEnvConfig().HTTP_PORT)
	err := restAdapter.Adapter.Start(utils.GetEnvConfig().HTTP_PORT)

	if err != nil {
		log.Fatal(err)
	}
}