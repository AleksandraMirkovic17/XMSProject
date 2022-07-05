package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       LoadEnvVariable("USER_SERVICE_PORT"),
		UserDBHost: LoadEnvVariable("MONGO_DB_HOST"),
		UserDBPort: LoadEnvVariable("MONGO_DB_PORT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
