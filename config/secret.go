package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ADMINGRPCPORT string
	ADMINUSERNAME string
	ADMINPASSWORD string
	ENVIRONMENT   string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}

	return Config{
		ENVIRONMENT:   os.Getenv("ENVIRONMENT"),
		ADMINGRPCPORT: os.Getenv("ADMINGRPCPORT"),
		ADMINUSERNAME: os.Getenv("ADMINUSERNAME"),
		ADMINPASSWORD: os.Getenv("ADMINPASSWORD"),
	}
}
