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
	DistributeToConnections CreateNotificationCommandType = iota
	GenerateContent
	CreateNotification
	UnknownCommand
)

type CreateNotificationCommand struct {
	Notification NotificationDetails
	Type CreateNotificationCommandType
}

type CreateNotificationReplyType int8

const (
	DistributeToConnectionsSuccess CreateNotificationReplyType = iota
	GenerateContentSuccess
	CreateNotificationSuccess
	GenericFailure
	UnknownReply
)

type CreateNotificationReply struct {
	Notification NotificationDetails
	Type CreateNotificationReplyType
}
