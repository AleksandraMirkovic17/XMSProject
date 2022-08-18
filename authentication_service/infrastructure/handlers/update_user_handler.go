package handlers

import (
	"AuthenticationService/application"
	"AuthenticationService/infrastructure/mappers"
	"fmt"
	saga "github.com/dislinked/common/saga/messaging"
	events "github.com/dislinked/common/saga/update_order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserCommandHandler struct {
	service           *application.AuthenticationService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateUserCommandHandler(service *application.AuthenticationService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateUserCommandHandler, error) {
	o := &UpdateUserCommandHandler{
		service:           service,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *UpdateUserCommandHandler) handle(command *events.UpdateUserCommand) {
	println("Nalazim se u handleru authentication servisa za update")
	id, err := primitive.ObjectIDFromHex(command.User.Id)
	reply := events.UpdateUserReply{
		User:    command.User,
		UserOld: command.OldUser,
	}
	if err != nil {
		println(command.User.Id)
		println("nemoguce id kovertovati")
		println("Failed to update user  auth profile")
		reply.Type = events.AuthenticationServiceNotUpdated
		_ = handler.replyPublisher.Publish(reply)
		return
	}

	print("Pre switchovanja")
	print(command.Type)
	switch command.Type {
	case events.AuthenticationServiceUpdate:
		println("Update authentication servisa")
		_, err := handler.service.GetOne(id)
		if err != nil {
			println("User does not existi")
			reply.Type = events.AuthenticationServiceNotUpdated
			break
		}
		err = handler.service.Update(mappers.MapUpdateCommandToAuthUser(&command.User))
		if err != nil {
			println("Failed to update user profile")
			reply.Type = events.AuthenticationServiceNotUpdated
		}
		println("uspeso update authentication servis")
		reply.Type = events.AuthenticationServiceUpdated

	case events.RollbackAuthenticationService:
		fmt.Println("Auth service:Rollback update authentication servisa")
		_, err := handler.service.GetOne(id)
		if err != nil {
			println("User does not existi")
			reply.Type = events.UnknownReply
			break
		}
		err = handler.service.Update(mappers.MapUpdateCommandToAuthUser(&command.OldUser))
		if err != nil {
			println("Failed to rollback user profile")
			reply.Type = events.UnknownReply
			break
		}
		println("uspeso rollback authentication servis")
		reply.Type = events.AuthenticationServiceRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
