package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                       string
	AuthDBHost                 string
	AuthDBPort                 string
	ApiGatewayHost             string
	ApiGatewayPort             string
	NatsHost                   string
	NatsPort                   string
	NatsUser                   string
	NatsPass                   string
	RegisterUserCommandSubject string
	RegisterUserReplySubject   string
	UpdateUserCommandSubject   string
	UpdateUserReplySubject     string
}

func NewConfig() *Config {

	return &Config{
		Port:                       LoadEnvVariable("AUTHENTICATION_SERVICE_PORT"),
		AuthDBHost:                 LoadEnvVariable("MONGO_DB_HOST"),
		AuthDBPort:                 LoadEnvVariable("MONGO_DB_PORT"),
		ApiGatewayPort:             LoadEnvVariable("AUTHENTICATION_SERVICE_HOST"),
		ApiGatewayHost:             LoadEnvVariable("AUTHENTICATION_SERVICE_PORT"),
		NatsHost:                   LoadEnvVariable("NATS_HOST"),
		NatsPort:                   LoadEnvVariable("NATS_PORT"),
		NatsUser:                   LoadEnvVariable("NATS_USER"),
		NatsPass:                   LoadEnvVariable("NATS_PASS"),
		RegisterUserCommandSubject: LoadEnvVariable("REGISTER_USER_COMMAND_SUBJECT"),
		RegisterUserReplySubject:   LoadEnvVariable("REGISTER_USER_REPLY_SUBJECT"),
		UpdateUserReplySubject:     LoadEnvVariable("UPDATE_USER_REPLY_SUBJECT"),
		UpdateUserCommandSubject:   LoadEnvVariable("UPDATE_USER_COMMAND_SUBJECT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
