package handlers

import (
	"UserService/application"
	events "github.com/dislinked/common/saga/friend_posted_notification"
	saga "github.com/dislinked/common/saga/messaging"
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

func (handler *FriendPostedNotificationHandler) handle(command events.FriendPostNotificationCommand) {
	println("Nalazim se u hendleru user servisa za slanje notifikacija za objavljene postove prijatelja")
	print("Command type je: ")
	println(command.Type)
	reply := events.FriendPostNotificationReply{
		Notification: command.Notification,
	}

	switch command.Type {
	case events.GetConnectionsWithTurnedOnNotifications:
		println("Getting connections with turned on notification")
		//TO DO: dodati
		reply.Type = events.JustConnectionsNotificationTurnedOnSuccess
		break
	default:
		reply.Type = events.UnknownReply

	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
