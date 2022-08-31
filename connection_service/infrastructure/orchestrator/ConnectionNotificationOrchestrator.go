package orchestrator

import (
	events "github.com/dislinked/common/saga/connection_notification"
	saga "github.com/dislinked/common/saga/messaging"
)

type ConnectionNotificationOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewConnectionNotificationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*ConnectionNotificationOrchestrator, error) {
	o := &ConnectionNotificationOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *ConnectionNotificationOrchestrator) handle(reply *events.ConnectionNotificationReply) {
	println("Nalazim se u hendleru update orkestratora")
	command := events.ConnectionNotificationCommand{
		Notification: events.ConnectionNotification{
			Content:    reply.Notification.Content,
			SenderId:   reply.Notification.SenderId,
			ReceiverId: reply.Notification.ReceiverId,
			Sender:     reply.Notification.Sender,
			Receiver:   reply.Notification.Receiver,
			Request:    reply.Notification.Request,
		},
		Type: o.nextCommand(reply.Type),
	}
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
		println("Publisovana je komanda u connection orkestratoru")
	}

}

func (o *ConnectionNotificationOrchestrator) nextCommand(reply events.ConnectionNotificationReplyType) events.ConnectionNotificationCommandType {
	switch reply {
	case events.NotificationsAreTurnedOff:
		println("Notification are turned of for user")
		return events.CancelSendingNotification
	case events.ReceiverFilled:
		println("Informations about receiver are collected, send notification")
		return events.SendNotification
	case events.NotificationSent:
		println("Notification is sent!")
		return events.UnknownCommand
	case events.NotificationNotSent:
		println("Notification is not sent, due some errors in notification service!")
		return events.UnknownCommand
	default:
		return events.UnknownCommand
	}
}

func (o *ConnectionNotificationOrchestrator) Start(notification *events.ConnectionNotification) error {
	println("Startovanje orkestratora u connections servisu")
	event := &events.ConnectionNotificationCommand{
		Notification: *notification,
		Type:         events.CheckAndGetReceiverById}

	return o.commandPublisher.Publish(event)

}
