package handlers

import (
	"UserService/application"
	"UserService/infrastructure/mappers"
	saga "github.com/dislinked/common/saga/messaging"
	events "github.com/dislinked/common/saga/update_order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateUserCommandHandler(service *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateUserCommandHandler, error) {
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
	reply := events.UpdateUserReply{
		User:    command.User,
		UserOld: command.OldUser,
	}
	switch command.Type {
	case events.RollebackUserProfile:
		println("Rolling back user profile")
		id, err := primitive.ObjectIDFromHex(command.User.Id)
		if err != nil {
			println("Error convert id")
			return
		}

		_, err = handler.userService.GetOne(id)
		if err != nil {
			println("User doe not exists!")
			return
		}

		err = handler.userService.Update(id, mappers.MapEventUserToDomainUser(command.OldUser))
		if err != nil {
			println("Failed to rollback user profile")
			return
		}
		reply.Type = events.UserProfileRolledBack
		break
	case events.UserProfileUpdate:
		println("Updating user profile")
		id, err := primitive.ObjectIDFromHex(command.User.Id)
		if err != nil {
			println("Error convert id")
			return
		}

		_, err = handler.userService.GetOne(id)
		if err != nil {
			println("User doe not exists!")
			return
		}

		err = handler.userService.Update(id, mappers.MapEventUserToDomainUser(command.User))
		if err != nil {
			println("Failed to update user profile")
			return
		}
		reply.Type = events.UserProfileUpdated
		break
	default:
		reply.Type = events.UnknownReply

	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
