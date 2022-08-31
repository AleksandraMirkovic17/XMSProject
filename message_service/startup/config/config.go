package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                              string
	MessageDBHost                     string
	MessageDBPort                     string
	NatsHost                          string
	NatsPort                          string
	NatsUser                          string
	NatsPass                          string
	MessageNotificationCommandSubject string
	MessageNotificationReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                              LoadEnvVariable("MESSAGE_SERVICE_PORT"),
		MessageDBHost:                     LoadEnvVariable("MONGO_DB_HOST"),
		MessageDBPort:                     LoadEnvVariable("MONGO_DB_PORT"),
		NatsHost:                          LoadEnvVariable("NATS_HOST"),
		NatsPort:                          LoadEnvVariable("NATS_PORT"),
		NatsUser:                          LoadEnvVariable("NATS_USER"),
		NatsPass:                          LoadEnvVariable("NATS_PASS"),
		MessageNotificationCommandSubject: LoadEnvVariable("MESSAGE_NOTIFICATION_COMMAND_SUBJECT"),
		MessageNotificationReplySubject:   LoadEnvVariable("MESSAGE_NOTIFICATION_REPLY_SUBJECT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
