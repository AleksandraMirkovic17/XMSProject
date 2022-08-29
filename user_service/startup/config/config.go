package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port                       string
	UserDBHost                 string
	UserDBPort                 string
	NatsHost                   string
	NatsPort                   string
	NatsUser                   string
	NatsPass                   string
	RegisterUserCommandSubject string
	RegisterUserReplySubject   string
	UpdateUserCommandSubject   string
	UpdateUserReplySubject     string
	FriendPostedCommandSubject string
	FriendPostedReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                       LoadEnvVariable("USER_SERVICE_PORT"),
		UserDBHost:                 LoadEnvVariable("MONGO_DB_HOST"),
		UserDBPort:                 LoadEnvVariable("MONGO_DB_PORT"),
		NatsHost:                   LoadEnvVariable("NATS_HOST"),
		NatsPort:                   LoadEnvVariable("NATS_PORT"),
		NatsUser:                   LoadEnvVariable("NATS_USER"),
		NatsPass:                   LoadEnvVariable("NATS_PASS"),
		RegisterUserCommandSubject: LoadEnvVariable("REGISTER_USER_COMMAND_SUBJECT"),
		RegisterUserReplySubject:   LoadEnvVariable("REGISTER_USER_REPLY_SUBJECT"),
		UpdateUserReplySubject:     LoadEnvVariable("UPDATE_USER_REPLY_SUBJECT"),
		UpdateUserCommandSubject:   LoadEnvVariable("UPDATE_USER_COMMAND_SUBJECT"),
		FriendPostedCommandSubject: LoadEnvVariable("CREATE_NOTIFICATION_COMMAND_SUBJECT"),
		FriendPostedReplySubject:   LoadEnvVariable("CREATE_NOTIFICATION_REPLY_SUBJECT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
