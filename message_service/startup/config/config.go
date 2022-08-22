package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	MessageDBHost string
	MessageDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          LoadEnvVariable("MESSAGE_SERVICE_PORT"),
		MessageDBHost: LoadEnvVariable("MONGO_DB_HOST"),
		MessageDBPort: LoadEnvVariable("MONGO_DB_PORT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
