package orchestrators

import (
	events "github.com/dislinked/common/saga/create_notification"
	saga "github.com/dislinked/common/saga/messaging"
)

type FriendPostedNotificationOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewFriendPostedNotificationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*FriendPostedNotificationOrchestrator, error) {
	o := &FriendPostedNotificationOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *FriendPostedNotificationOrchestrator) handle(reply *events.CreateNotificationReply) {
	println("Nalazim se u hendleru friend posted orkestratora")
	command := events.CreateNotificationCommand{
		Notification: events.NotificationDetails{
			User:    reply.Notification.User,
			Content: reply.Notification.Content,
			Url:     reply.Notification.Url,
			Seen:    reply.Notification.Seen,
		},
		Type: o.nextCommand(reply.Type),
	}
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
		print("Publisovana je komanda ")
		println(command.Type)
	}

}

func (o *FriendPostedNotificationOrchestrator) nextCommand(reply events.CreateNotificationReplyType) events.CreateNotificationCommandType {
	/*switch reply {
	case events.ConnectionsFail:
		println("It was imposible to get sender connections!")
		return events.RollbackPostService
	case events.ConnectionsSuccess:
		println("Successfully collected all senders connections!")
		return events.GetConnectionsWithTurnedOnNotifications
	case events.JustConnectionsNotificationTurnedOnSuccess:
		println("Successfully got just connections with turned on notifications!")
		return events.SendNotification
	case events.JusConnectionsNotificationTurnedOnFailed:
		return events.RollbackPostService
	case events.NotificationSent:
		println("Notification for friend post successfully sent!")
		return events.SuccessfulyPosted
	case events.NotificationNotSent:
		return events.RollbackPostService
	case events.PostServiceRolledBack:
		println("Post service rolled back")
		return events.FiledToPost
	default:
		return events.UnknownCommand

	}*/
	return events.UnknownCommand
}

func (o *FriendPostedNotificationOrchestrator) Start(notification *events.NotificationDetails) error {
	println("Starting orchestrator for sending notifications for posts!")
	event := events.CreateNotificationCommand{
		Type:         events.GetConnections,
		Notification: *notification,
	}
	return o.commandPublisher.Publish(event)
}
