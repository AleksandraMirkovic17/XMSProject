package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                                 string
	NotificationDBHost                   string
	NotificationDBPort                   string
	ApiGatewayHost                       string
	ApiGatewayPort                       string
	NatsHost                             string
	NatsPort                             string
	NatsUser                             string
	NatsPass                             string
	FriendPostedCommandSubject           string
	FriendPostedReplySubject             string
	ConnectionNotificationCommandSubject string
	ConnectionNotificationReplySubject   string
	MessageNotificationCommandSubject    string
	MessageNotificationReplySubject      string
}

func NewConfig() *Config {

	return &Config{
		Port:                                 LoadEnvVariable("NOTIFICATION_SERVICE_PORT"),
		NotificationDBHost:                   LoadEnvVariable("MONGO_DB_HOST"),
		NotificationDBPort:                   LoadEnvVariable("MONGO_DB_PORT"),
		NatsHost:                             LoadEnvVariable("NATS_HOST"),
		NatsPort:                             LoadEnvVariable("NATS_PORT"),
		NatsUser:                             LoadEnvVariable("NATS_USER"),
		NatsPass:                             LoadEnvVariable("NATS_PASS"),
		FriendPostedCommandSubject:           LoadEnvVariable("CREATE_NOTIFICATION_COMMAND_SUBJECT"),
		FriendPostedReplySubject:             LoadEnvVariable("CREATE_NOTIFICATION_REPLY_SUBJECT"),
		ConnectionNotificationCommandSubject: LoadEnvVariable("CONNECTION_NOTIFICATION_COMMAND_SUBJECT"),
		ConnectionNotificationReplySubject:   LoadEnvVariable("CONNECTION_NOTIFICATION_REPLY_SUBJECT"),
		MessageNotificationCommandSubject:    LoadEnvVariable("MESSAGE_NOTIFICATION_COMMAND_SUBJECT"),
		MessageNotificationReplySubject:      LoadEnvVariable("MESSAGE_NOTIFICATION_REPLY_SUBJECT"),
	}
}

func LoadEnvVariable(key string) string {
	godotenv.Load("../.env")
	return os.Getenv(key)
}
