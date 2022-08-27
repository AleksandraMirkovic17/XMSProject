package friend_posted_notification

type FriendPostNotification struct {
	postID                string
	Content               string
	RedirectPath          string
	NotificationSender    string
	NotificationReceivers []string
}

type FriendPostNotificationCommandType int8

const (
	GetConnections FriendPostNotificationCommandType = iota
	GetConnectionsWithTurnedOnNotifications
	RollbackPostService
	SendNotification
	SuccessfulyPosted
	FiledToPost
	UnknownCommand
)

type FriendPostNotificationReplyType int8

const (
	ConnectionsSuccess FriendPostNotificationReplyType = iota
	ConnectionsFail
	JustConnectionsNotificationTurnedOnSuccess
	JusConnectionsNotificationTurnedOnFailed
	NotificationSent
	NotificationNotSent
	PostServiceRolledBack
	UnknownReply
)

type FriendPostNotificationCommand struct {
	Notification FriendPostNotification
	Type         FriendPostNotificationCommandType
}

type FriendPostNotificationReply struct {
	Notification FriendPostNotification
	Type         FriendPostNotificationReplyType
}
