package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                       string
	PostDBHost                 string
	PostDBPort                 string
	NatsHost                   string
	NatsPort                   string
	NatsUser                   string
	NatsPass                   string
	FriendPostedCommandSubject string
	FriendPostedReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                       LoadEnvVariable("POST_SERVICE_PORT"),
		PostDBHost:                 LoadEnvVariable("MONGO_DB_HOST"),
		PostDBPort:                 LoadEnvVariable("MONGO_DB_PORT"),
		NatsHost:                   LoadEnvVariable("NATS_HOST"),
		NatsPort:                   LoadEnvVariable("NATS_PORT"),
		NatsUser:                   LoadEnvVariable("NATS_USER"),
		NatsPass:                   LoadEnvVariable("NATS_PASS"),
		FriendPostedCommandSubject: LoadEnvVariable("CREATE_NOTIFICATION_COMMAND_SUBJECT"),
		FriendPostedReplySubject:   LoadEnvVariable("CREATE_NOTIFICATION_REPLY_SUBJECT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
