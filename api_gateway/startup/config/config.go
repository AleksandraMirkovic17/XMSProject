package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port     string
	PostHost string
	PostPort string
	UserHost string
	UserPort string
	AuthHost string
	AuthPort string
}

func NewConfig() *Config {
	return &Config{
		Port:     LoadEnvVariable("GATEWAY_PORT"),
		PostHost: LoadEnvVariable("POST_SERVICE_HOST"),
		PostPort: LoadEnvVariable("POST_SERVICE_PORT"),
		UserHost: LoadEnvVariable("USER_SERVICE_HOST"),
		UserPort: LoadEnvVariable("USER_SERVICE_PORT"),
		AuthHost: LoadEnvVariable("AUTHENTICATION_SERVICE_HOST"),
		AuthPort: LoadEnvVariable("AUTHENTICATION_SERVICE_PORT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
