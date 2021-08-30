package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	AppName     string
	AppPort     int
	LogLevel    string
	Environment string
	JWTSecret   string
	DbUser      string
	DbPassword  string
	DbAddr      string
	DbPort      string
	DbName      string
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Fatal("failed load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig := &Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetInt("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
		JWTSecret:   GetString("JWT_SECRET"),
		DbUser:      GetString("DB_USER"),
		DbPassword:  GetString("DB_PASSWORD"),
		DbAddr:      GetString("DB_ADDR"),
		DbPort:      GetString("DB_PORT"),
		DbName:      GetString("DB_NAME"),
	}

	return appConfig
}
