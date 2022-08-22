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
	ConnHost string
	ConnPort string
	JobHost  string
	JobPort  string
	MessageHost  string
	MessagePort  string
}

func NewConfig() *Config {
	return &Config{
		Port:     		LoadEnvVariable("GATEWAY_PORT"),
		PostHost: 		LoadEnvVariable("POST_SERVICE_HOST"),
		PostPort: 		LoadEnvVariable("POST_SERVICE_PORT"),
		UserHost: 		LoadEnvVariable("USER_SERVICE_HOST"),
		UserPort: 		LoadEnvVariable("USER_SERVICE_PORT"),
		AuthHost: 		LoadEnvVariable("AUTHENTICATION_SERVICE_HOST"),
		AuthPort: 		LoadEnvVariable("AUTHENTICATION_SERVICE_PORT"),
		ConnHost: 		LoadEnvVariable("CONNECTION_SERVICE_HOST"),
		ConnPort: 		LoadEnvVariable("CONNECTION_SERVICE_PORT"),
		JobHost:  		LoadEnvVariable("JOB_SERVICE_HOST"),
		JobPort:  		LoadEnvVariable("JOB_SERVICE_PORT"),
		MessageHost:  	LoadEnvVariable("MESSAGE_SERVICE_HOST"),
		MessagePort:  	LoadEnvVariable("MESSAGE_SERVICE_PORT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
