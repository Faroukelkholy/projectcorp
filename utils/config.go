package utils

import (
	"github.com/JeremyLoy/config"
	"github.com/joho/godotenv"
	"log"
)

type EnvConfig struct {
	HTTP_SERVER string
	HTTP_PORT   string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	EMPLOYEE_Domain string
	GETEMPLOYEES_URL string
}

var ec EnvConfig

func GetEnvConfig() *EnvConfig {
	errGodot := godotenv.Load()
	if errGodot != nil {
		log.Fatal("Error loading .env file")
	}
	err := config.FromEnv().To(&ec)
	if err != nil {
		log.Fatal(err)
	}
	return &ec
}
