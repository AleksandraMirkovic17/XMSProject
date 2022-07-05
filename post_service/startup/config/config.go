package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port       string
	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       LoadEnvVariable("POST_SERVICE_PORT"),
		PostDBHost: LoadEnvVariable("MONGO_DB_HOST"),
		PostDBPort: LoadEnvVariable("MONGO_DB_PORT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
