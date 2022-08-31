package handler

import (
	"MessageService/application"
	events "github.com/dislinked/common/saga/message_notification"
	saga "github.com/dislinked/common/saga/messaging"
)

type MessageNotificationHandler struct {
	service           *application.MessageService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewMessageNotificatioHandler(service *application.MessageService, publiser saga.Publisher, subscriber saga.Subscriber) (*MessageNotificationHandler, error) {
	o := &MessageNotificationHandler{
		service:           service,
		replyPublisher:    publiser,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *MessageNotificationHandler) handle(command events.MessageNotificationCommand) {
	println("Nalazim se u hendleru message servisa za slanje notifikacija za poruke")
	print("Command type je: ")
	println(command.Type)
	reply := events.MessageNotificationReply{
		Notification: command.Notification,
	}

	switch command.Type {
	case events.CancelSendingNotification:
		println("Sending notification is cancelled")
		reply.Type = events.UnknownReply
		break
	default:
		reply.Type = events.UnknownReply

	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
