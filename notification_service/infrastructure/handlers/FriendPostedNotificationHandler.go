package handlers

import (
	"NotificationService/application"
	"NotificationService/domain"
	"time"

	events "github.com/dislinked/common/saga/create_notification"
	saga "github.com/dislinked/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendPostedNotificationHandler struct {
	service           *application.NotificationService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewFriendPostedNotificationHandler(service *application.NotificationService, publiser saga.Publisher, subscriber saga.Subscriber) (*FriendPostedNotificationHandler, error) {
	o := &FriendPostedNotificationHandler{
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

func mapNotificationCommandToDomain(notification *events.NotificationDetails) *domain.Notification {
	userDomain := &domain.Notification{
		Id:      primitive.NewObjectID(),
		User:    notification.User,
		Content: notification.Content,
		Url:     notification.Url,
		Seen:    notification.Seen,
		Date:    time.Now(),
	}
	return userDomain
}

func (handler *FriendPostedNotificationHandler) handle(command events.CreateNotificationCommand) {
	reply := events.CreateNotificationReply{
		Notification: command.Notification,
	}

	switch command.Type {
	case events.CreateNotification:
		println("Creating notification for user: " + command.Notification.User)
		handler.service.Insert(mapNotificationCommandToDomain(&command.Notification))
		reply.Type = events.CreateNotificationSuccess
		handler.replyPublisher.Publish(reply)
		break
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		handler.replyPublisher.Publish(reply)
	}
}
