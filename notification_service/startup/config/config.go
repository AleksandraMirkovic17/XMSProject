package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                             string
	NotificationDBHost               string
	NotificationDBPort               string
	ApiGatewayHost                   string
	ApiGatewayPort                   string
	NatsHost                         string
	NatsPort                         string
	NatsUser                         string
	NatsPass                         string
	CreateNotificationCommandSubject string
	CreateNotificationReplySubject   string
}

func NewConfig() *Config {

	return &Config{
		Port:                             LoadEnvVariable("NOTIFICATION_SERVICE_PORT"),
		NotificationDBHost:               LoadEnvVariable("MONGO_DB_HOST"),
		NotificationDBPort:               LoadEnvVariable("MONGO_DB_PORT"),
		ApiGatewayPort:                   LoadEnvVariable("NOTIFICATION_SERVICE_HOST"),
		ApiGatewayHost:                   LoadEnvVariable("NOTIFICATION_SERVICE_PORT"),
		NatsHost:                         LoadEnvVariable("NATS_HOST"),
		NatsPort:                         LoadEnvVariable("NATS_PORT"),
		NatsUser:                         LoadEnvVariable("NATS_USER"),
		NatsPass:                         LoadEnvVariable("NATS_PASS"),
		CreateNotificationCommandSubject: LoadEnvVariable("CREATE_NOTIFICATION_COMMAND_SUBJECT"),
		CreateNotificationReplySubject:   LoadEnvVariable("CREATE_NOTIFICATION_REPLY_SUBJECT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
