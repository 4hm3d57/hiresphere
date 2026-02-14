package config


import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	HTTPport string
	AppEnv string
	DBUrl string
}


func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	
	return &Config {
		HTTPport: os.Getenv("HTTP_PORT"),
		AppEnv: os.Getenv("APP_ENV"),
		DBUrl: os.Getenv("DB_URL"),
	}

}
