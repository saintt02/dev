package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Cargar el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file: ", err)
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
