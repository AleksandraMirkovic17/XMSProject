package create_notification

type NotificationDetails struct {
	string id = 1;
	string user = 2;
	string content = 3;
	string url = 4;
	bool seen = 5;
	google.protobuf.Timestamp date = 6;
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
