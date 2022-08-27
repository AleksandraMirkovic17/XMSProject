package handlers

import (
	"PostService/application"
	events "github.com/dislinked/common/saga/friend_posted_notification"
	saga "github.com/dislinked/common/saga/messaging"
)

type FriendPostedNotificationHandler struct {
	service           *application.PostService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewFriendPostedNotificationHandler(service *application.PostService, publiser saga.Publisher, subsciber *saga.Subscriber) (*FriendPostedNotificationHandler, error) {
	o := &FriendPostedNotificationHandler{
		service:           service,
		replyPublisher:    publiser,
		commandSubscriber: subsciber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *FriendPostedNotificationHandler) handle(command events.FriendPostNotificationCommand) {
	println("Nalazim se u hendleru post servisa za slanje notifikacija za objavljene postove prijatelja")
	print("Command type je: ")
	println(command.Type)
	reply := events.FriendPostNotificationReply{
		Notification: command.Notification,
	}

	switch command.Type {
	case events.FiledToPost:
		println("Failed to post")
		reply.Type = events.UnknownReply
		break
	case events.RollbackPostService:
		println("Have to rollback post service")
		/*err := handler.service.Delete(command.Notification.postID)*/
		reply.Type = events.PostServiceRolledBack
		break
	case events.SuccessfulyPosted:
		reply.Type = events.UnknownReply
		break
	default:
		reply.Type = events.UnknownReply

	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher
	}
}
