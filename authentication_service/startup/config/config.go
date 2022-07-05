package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port           string
	AuthDBHost     string
	AuthDBPort     string
	ApiGatewayHost string
	ApiGatewayPort string
}

func NewConfig() *Config {

	return &Config{
		Port:           LoadEnvVariable("AUTHENTICATION_SERVICE_PORT"),
		AuthDBHost:     LoadEnvVariable("MONGO_DB_HOST"),
		AuthDBPort:     LoadEnvVariable("MONGO_DB_PORT"),
		ApiGatewayPort: LoadEnvVariable("AUTHENTICATION_SERVICE_HOST"),
		ApiGatewayHost: LoadEnvVariable("AUTHENTICATION_SERVICE_PORT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
