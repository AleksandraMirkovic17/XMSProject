package api

import (
	"ConnectionService/application"
	events "github.com/dislinked/common/saga/connection_notification"
	saga "github.com/dislinked/common/saga/messaging"
)

type ConnectionNotificationHandler struct {
	service           *application.ConnectionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewConnectionNotificatioHandler(service *application.ConnectionService, publiser saga.Publisher, subscriber saga.Subscriber) (*ConnectionNotificationHandler, error) {
	o := &ConnectionNotificationHandler{
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

func (handler *ConnectionNotificationHandler) handle(command events.ConnectionNotificationCommand) {
	println("Nalazim se u hendleru connection servisa za slanje notifikacija za konekcije")
	print("Command type je: ")
	println(command.Type)
	reply := events.ConnectionNotificationReply{
		Notification: command.Notification,
	}

	switch command.Type {
	case events.CancelSendingNotification:
		println("Sending notification is cancelled")
		//TO DO: dodati
		reply.Type = events.UnknownReply
		break
	default:
		reply.Type = events.UnknownReply

	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
