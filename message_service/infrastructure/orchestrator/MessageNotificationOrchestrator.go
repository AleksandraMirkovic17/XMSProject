package orchestrator

import (
	events "github.com/dislinked/common/saga/message_notification"
	saga "github.com/dislinked/common/saga/messaging"
)

type MessageNotificationOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewMessageNotificationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*MessageNotificationOrchestrator, error) {
	o := &MessageNotificationOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *MessageNotificationOrchestrator) handle(reply *events.MessageNotificationReply) {
	println("Nalazim se u hendleru message notification orkestratora")
	command := events.MessageNotificationCommand{
		Notification: events.MessageNotification{
			Content:    reply.Notification.Content,
			SenderId:   reply.Notification.SenderId,
			ReceiverId: reply.Notification.ReceiverId,
			Sender:     reply.Notification.Sender,
			Receiver:   reply.Notification.Receiver,
		},
		Type: o.nextCommand(reply.Type),
	}
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
		println("Publisovana je komanda u connection orkestratoru")
	}

}

func (o *MessageNotificationOrchestrator) nextCommand(reply events.MessageNotificationReplyType) events.MessageNotificationCommandType {
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

func (o *MessageNotificationOrchestrator) Start(notification *events.MessageNotification) error {
	println("Startovanje orkestratora u message servisu")
	event := &events.MessageNotificationCommand{
		Notification: *notification,
		Type:         events.CheckAndGetReceiverById}

	return o.commandPublisher.Publish(event)

}
