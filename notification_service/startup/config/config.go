package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	NotificationDBHost string
	NotificationDBPort string

	UserHost string
	UserPort string

	ConnectionHost string
	ConnectionPort string
}

func NewConfig() *Config {
	return &Config{
		Port:               LoadEnvVariable("NOTIFICATION_SERVICE_PORT"),
		NotificationDBHost: LoadEnvVariable("MONGO_DB_HOST"),
		NotificationDBPort: LoadEnvVariable("MONGO_DB_PORT"),

		UserHost: LoadEnvVariable("USER_SERVICE_HOST"),
		UserPort: LoadEnvVariable("USER_SERVICE_PORT"),

		ConnectionHost: LoadEnvVariable("CONNECTION_SERVICE_HOST"),
		ConnectionPort: LoadEnvVariable("CONNECTION_SERVICE_PORT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
