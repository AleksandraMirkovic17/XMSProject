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
