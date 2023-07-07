package configs

import (
	"os"
)

var config *appConfig

type appConfig struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() error {
	config = new(appConfig)
	config.API = APIConfig{
		Port: os.Getenv("HTTP_PORT"),
	}
	config.DB = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	return nil
}

func GetDB() DBConfig {
	return config.DB
}

func GetApiPort() string {
	return config.API.Port
}
