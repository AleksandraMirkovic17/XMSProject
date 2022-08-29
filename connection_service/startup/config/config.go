package config

import (
	"os"

	"github.com/joho/godotenv"
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

	NatsHost                             string
	NatsPort                             string
	NatsUser                             string
	NatsPass                             string
	RegisterUserCommandSubject           string
	RegisterUserReplySubject             string
	ConnectionNotificationCommandSubject string
	ConnectionNotificationReplySubject   string
	UpdateUserCommandSubject             string
	UpdateUserReplySubject               string
	FriendPostedCommandSubject           string
	FriendPostedReplySubject             string
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
