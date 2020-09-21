package main

import (
	"log"
	"projectcorp/utils"
)

func main(){
	log.Println("portttt:", utils.GetEnvConfig().HTTP_PORT)
	log.Println("GETEMPLOYEES_URL:", utils.GetEnvConfig().GETEMPLOYEES_URL)
}