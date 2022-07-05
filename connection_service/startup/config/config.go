package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port          string
	Host          string
	Neo4jUri      string
	Neo4jHost     string
	Neo4jPort     string
	Neo4jUsername string
	Neo4jPassword string

	UserHost string
	UserPort string
}

func NewConfig() *Config {
	return &Config{
		Host: LoadEnvVariable("CONNECTION_SERVICE_HOST"),
		Port: LoadEnvVariable("CONNECTION_SERVICE_PORT"),

		Neo4jUri:      LoadEnvVariable("NEO4J_URI"),
		Neo4jHost:     LoadEnvVariable("NEO4J_HOST"),
		Neo4jPort:     LoadEnvVariable("NEO4J_PORT"),
		Neo4jUsername: LoadEnvVariable("NEO4J_USERNAME"),
		Neo4jPassword: LoadEnvVariable("NEO4J_PASSWORD"),

		UserHost: LoadEnvVariable("USER_SERVICE_HOST"),
		UserPort: LoadEnvVariable("USER_SERVICE_PORT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
