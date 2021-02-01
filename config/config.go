package config

import (
	"log"

	"github.com/JeremyLoy/config"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServer      string
	HTTPPort        string
	DBHost          string
	DBPort          int
	DBUser          string
	DBPASS          string
	DBName          string
	EmployeDomain   string
	GetEmployeesURL string
}

var cfg Config

func Parse() *Config {
	errGodot := godotenv.Load()
	if errGodot != nil {
		log.Fatal("Error loading .env file")
	}
	err := config.FromEnv().To(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
