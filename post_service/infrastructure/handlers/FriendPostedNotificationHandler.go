package handlers

import (
	"PostService/application"

	events "github.com/dislinked/common/saga/create_notification"
	saga "github.com/dislinked/common/saga/messaging"
)

type FriendPostedNotificationHandler struct {
	service           *application.PostService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewFriendPostedNotificationHandler(service *application.PostService, publiser saga.Publisher, subscriber *saga.Subscriber) (*FriendPostedNotificationHandler, error) {
	o := &FriendPostedNotificationHandler{
		service:           service,
		replyPublisher:    publiser,
		commandSubscriber: *subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *FriendPostedNotificationHandler) handle(command events.CreateNotificationCommand) {
	println("Nalazim se u hendleru post servisa za slanje notifikacija za objavljene postove prijatelja")
	print("Command type je: ")
	println(command.Type)
	reply := events.CreateNotificationReply{
		Notification: command.Notification,
	}

	/*switch command.Type {
	case events.CreateNotification:
		println("Failed to post")
		reply.Type = events.UnknownReply
		break
	case events.RollbackPostService:
		println("Have to rollback post service")
		reply.Type = events.PostServiceRolledBack
		break
	case events.SuccessfulyPosted:
		reply.Type = events.UnknownReply
		break
	default:
		reply.Type = events.UnknownReply

	}*/
	reply.Type = events.CreatedNotification
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher
	}
}
