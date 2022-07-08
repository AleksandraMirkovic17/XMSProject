package api

import (
	"UserService/application"
	"fmt"
	events "github.com/dislinked/common/saga/create_order"
	saga "github.com/dislinked/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(orderService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateOrderCommandHandler, error) {
	o := &CreateOrderCommandHandler{
		userService:       orderService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateOrderCommandHandler) handle(command *events.RegisterUserCommand) {

	order := mapCommandToUser(command) //ovde se postavlja id

	reply := events.RegisterUserReply{Order: command.Order}

	switch command.Type {
	case events.UserProfileCreate:
		err := handler.userService.Register(order)
		if err != nil {
			reply.Type = events.UserProfileNotCreated
			return
		}
		reply.Type = events.UserProfileCreated

	case events.RollebackUserProfile:
		fmt.Println("Auth service:Rollback authentication servisa")
		id, err := primitive.ObjectIDFromHex(command.Order.Id)
		if err != nil {
			return
		}
		toDelete, _ := handler.userService.GetOne(id)
		handler.userService.Delete(toDelete)
		reply.Type = events.UserProfileRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
