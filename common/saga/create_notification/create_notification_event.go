package create_notification

type NotificationDetails struct {
	Id 		string
	User 	string
	Content string
	Url		string
	Seen	bool
}

type CreateNotificationCommandType int8

const (
	GetConnections CreateNotificationCommandType = iota
	CreateNotification
	NotificationCreate
	UnknownCommand
)

type CreateNotificationCommand struct {
	Notification NotificationDetails
	Type CreateNotificationCommandType
}

type CreateNotificationReplyType int8

const (
	ConnectionsSuccess CreateNotificationReplyType = iota
	ConnectionsFail
	CreatedNotification
	NotificationCreated
	UnknownReply
)

type CreateNotificationReply struct {
	Notification NotificationDetails
	Type CreateNotificationReplyType
}
