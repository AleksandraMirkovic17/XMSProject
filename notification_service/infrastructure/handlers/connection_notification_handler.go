package handlers

import (
	"NotificationService/application"
	"NotificationService/domain"
	events "github.com/dislinked/common/saga/connection_notification"
	saga "github.com/dislinked/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ConnectionNotificationHandler struct {
	service           *application.NotificationService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewConnectionNotificatioHandler(service *application.NotificationService, publiser saga.Publisher, subscriber saga.Subscriber) (*ConnectionNotificationHandler, error) {
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
	println("Nalazim se u hendleru notification servisa za slanje notifikacija za konekcije")
	print("Command type je: ")
	println(command.Type)
	println("sender id")
	println(command.Notification.SenderId)
	println("Sender")
	println(command.Notification.Sender)
	reply := events.ConnectionNotificationReply{
		Notification: command.Notification,
	}

	switch command.Type {
	case events.SendNotification:
		println("Sending notification in notification service")
		//TO DO: dodati
		err := handler.service.Insert(&domain.Notification{
			Id:      primitive.NewObjectID(),
			User:    command.Notification.ReceiverId,
			Content: command.Notification.Content,
			Url:     "",
			Seen:    false,
			Date:    time.Time{},
		})
		if err != nil {
			println("Notification is not sent:", err.Error())
			reply.Type = events.NotificationNotSent
			break
		}
		reply.Type = events.NotificationSent
		break
	default:
		reply.Type = events.UnknownReply

	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
