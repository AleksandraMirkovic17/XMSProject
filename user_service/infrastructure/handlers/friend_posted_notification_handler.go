package handlers

import (
	"UserService/application"

	events "github.com/dislinked/common/saga/create_notification"
	saga "github.com/dislinked/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendPostedNotificationHandler struct {
	service           *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewFriendPostedNotificationHandler(service *application.UserService, publiser saga.Publisher, subscriber saga.Subscriber) (*FriendPostedNotificationHandler, error) {
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

func (handler *FriendPostedNotificationHandler) handle(command events.CreateNotificationCommand) {
	reply := events.CreateNotificationReply{
		Notification: command.Notification,
	}

	switch command.Type {
	case events.GenerateContent:
		userId, _ := primitive.ObjectIDFromHex(command.Notification.User)
		user, _ := handler.service.GetOne(userId)
		if user != nil {
			reply.Type = events.GenerateContentSuccess
			reply.Notification.Content = user.Username + " has made a new post"
			reply.Notification.Url = "profile/" + user.Id.Hex()
		} else {
			reply.Type = events.GenericFailure
		}
		break
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
