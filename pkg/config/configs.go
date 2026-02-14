package config


import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	HTTP_PORT string
	APP_ENV string
	DB_URL string
}


func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	
	return &Config {
		HTTP_PORT: os.Getenv("HTTP_PORT"),
		APP_ENV: os.Getenv("APP_ENV"),
		DB_URL: os.Getenv("DB_URL"),
	}

}
