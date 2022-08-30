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
		print("Next command: " + string(command.Type))
		o.commandPublisher.Publish(command)
	}

}

func (o *FriendPostedNotificationOrchestrator) nextCommand(reply events.CreateNotificationReplyType) events.CreateNotificationCommandType {
	switch reply {
	case events.DistributeToConnectionsSuccess:
		return events.GenerateContent
	case events.GenerateContentSuccess:
		return events.CreateNotification
	default:
		return events.UnknownCommand
	}
}

func (o *FriendPostedNotificationOrchestrator) Start(Notification *events.NotificationDetails) error {
	event := &events.CreateNotificationCommand{
		Notification: *Notification,
		Type:         events.CreateNotificationCommandType(events.DistributeToConnectionsSuccess),
	}
	return o.commandPublisher.Publish(event)
}
