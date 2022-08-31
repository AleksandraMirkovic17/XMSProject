package message_notification

type MessageNotification struct {
	Content    string
	SenderId   string
	ReceiverId string
	Sender     string
	Receiver   string
}

type MessageNotificationCommandType int8

const (
	CheckAndGetReceiverById MessageNotificationCommandType = iota
	SendNotification
	CancelSendingNotification
	UnknownCommand
)

type MessageNotificationReplyType int8

const (
	ReceiverFilled MessageNotificationReplyType = iota
	NotificationsAreTurnedOff
	NotificationSent
	NotificationNotSent
	UnknownReply
)

type MessageNotificationCommand struct {
	Notification MessageNotification
	Type         MessageNotificationCommandType
}

type MessageNotificationReply struct {
	Notification MessageNotification
	Type         MessageNotificationReplyType
}
