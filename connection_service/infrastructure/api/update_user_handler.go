package api

import (
	"ConnectionService/application"
	saga "github.com/dislinked/common/saga/messaging"
	events "github.com/dislinked/common/saga/update_order"
)

type UpdateUserCommandHandler struct {
	userService       *application.ConnectionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateUserCommandHandler(service *application.ConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateUserCommandHandler, error) {
	o := &UpdateUserCommandHandler{
		userService:       service,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *UpdateUserCommandHandler) handle(command events.UpdateUserCommand) {
	println("Nalazim se u handleru connection servisa za update")
	println(command.Type)
	reply := events.UpdateUserReply{
		User:    command.User,
		UserOld: command.OldUser,
	}
	print("Pre switchovanja: ")
	println(command.Type)
	switch command.Type {
	case events.UpdateUserNode:
		println("Updating user node")

		err, _ := handler.userService.UpdateUser(*mapEventUpdateUserToDomainUser(command.User))
		if err.Status != 200 {
			println("User doe not exists!")
			reply.Type = events.UserNodeFailedToUpdate
			break
		}
		reply.Type = events.UserNodeUpdated
		break
	case events.RollebackConnectionNode:
		println("Rollback user node")
		err, _ := handler.userService.UpdateUser(*mapEventUpdateUserToDomainUser(command.OldUser))
		if err.Status != 200 {
			println("User does not exists!")
			reply.Type = events.UserNodeFailedToUpdate
			break
		}
		reply.Type = events.ConnectionsRolledBack
		break
	default:
		reply.Type = events.UnknownReply

	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
