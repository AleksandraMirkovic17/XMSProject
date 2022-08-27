package create_notification

import "google.golang.org/grpc/credentials/google"

type NotificationDetails struct {
	Id 		string
	User 	string
	Content string
	Url		string
	Seen	bool
	Date 	google.protobuf.Timestamp
}

type CreateNotificationCommandType int8

const (
	CreateNotification CreateNotificationCommandType = iota
	NotificationCreate
	UnknownCommand
)

type CreateNotificationCommand struct {
	Notification NotificationDetails
	Type CreateNotificationCommandType
}

type CreateNotificationReplyType int8

const (
	CreatedNotification CreateNotificationReplyType = iota
	NotificationCreated
	UnknownReply
)

type CreateNotificationReply struct {
	Notification NotificationDetails
	Type CreateNotificationReplyType
}
