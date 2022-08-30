package api

import (
	"ConnectionService/application"

	events "github.com/dislinked/common/saga/create_notification"
	saga "github.com/dislinked/common/saga/messaging"
)

type FriendPostedNotificationHandler struct {
	service           *application.ConnectionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewFriendPostedNotificationHandler(service *application.ConnectionService, publiser saga.Publisher, subscriber saga.Subscriber) (*FriendPostedNotificationHandler, error) {
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
	case events.DistributeToConnections:
		friends, err := handler.service.GetFriends(command.Notification.User)
		if err != nil {
			reply.Type = events.UnknownReply
		} else {
			reply.Type = events.DistributeToConnectionsSuccess
			for _, friend := range friends {
				reply.Notification.User = friend.UserID
				handler.replyPublisher.Publish(reply)
			}
		}
		break
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply && reply.Type != events.DistributeToConnectionsSuccess {
		handler.replyPublisher.Publish(reply)
	}
}
