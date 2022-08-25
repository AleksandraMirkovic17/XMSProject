package handlers

import (
	"NotificationService/application"
	"NotificationService/domain"

	events "github.com/dislinked/common/saga/create_notification"
	saga "github.com/dislinked/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateNotificationCommandHandler struct {
	notificationService *application.NotificationService
	replyPublisher      saga.Publisher
	commandSubscriber   saga.Subscriber
}

func NewCreateNotificationCommandHandler(notificationService *application.NotificationService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateNotificationCommandHandler, error) {
	o := &CreateNotificationCommandHandler{
		notificationService: notificationService,
		replyPublisher:      publisher,
		commandSubscriber:   subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateNotificationCommandHandler) handle(command *events.CreateNotificationCommand) {
	id, err := primitive.ObjectIDFromHex(command.Notification.Id)
	if err != nil {
		return
	}
	notification := &domain.Notification{Id: id} //ovde se postavlja id
	handler.notificationService.Insert(notification)
	reply := events.CreateNotificationReply{Notification: command.Notification}
	reply.Type = events.NotificationCreated
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
