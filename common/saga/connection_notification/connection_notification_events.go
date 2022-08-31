package connection_notification

type ConnectionNotification struct {
	Content    string
	SenderId   string
	ReceiverId string
	Sender     string
	Receiver   string
	Request    bool
}

type ConnectionNotificationCommandType int8

const (
	CheckAndGetReceiverById ConnectionNotificationCommandType = iota
	SendNotification
	CancelSendingNotification
	UnknownCommand
)

type ConnectionNotificationReplyType int8

const (
	ReceiverFilled ConnectionNotificationReplyType = iota
	NotificationsAreTurnedOff
	NotificationSent
	NotificationNotSent
	UnknownReply
)

type ConnectionNotificationCommand struct {
	Notification ConnectionNotification
	Type         ConnectionNotificationCommandType
}

type ConnectionNotificationReply struct {
	Notification ConnectionNotification
	Type         ConnectionNotificationReplyType
}
