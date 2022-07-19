package api

import (
	"AuthenticationService/application"
	"AuthenticationService/domain"
	"fmt"

	events "github.com/dislinked/common/saga/create_order"
	saga "github.com/dislinked/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderCommandHandler struct {
	orderService      *application.AuthenticationService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(orderService *application.AuthenticationService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateOrderCommandHandler, error) {
	o := &CreateOrderCommandHandler{
		orderService:      orderService,
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
	id, err := primitive.ObjectIDFromHex(command.User.Id)
	if err != nil {
		return
	}
	order := &domain.User{ID: id} //ovde se postavlja id

	reply := events.RegisterUserReply{User: command.User}

	switch command.Type {
	case events.AuthenticationServiceUpdate:
		_, err := handler.orderService.Register(mapCommandToAuthUser(command))
		if err != nil {
			reply.Type = events.AuthenticationServiceNotUpdated
			return
		}
		reply.Type = events.AuthenticationServiceUpdated

	case events.RollbackAuthenticationService:
		fmt.Println("Auth service:Rollback authentication servisa")
		err := handler.orderService.DeleteById(order.ID)
		if err != nil {
			return
		}
		reply.Type = events.AuthenticationServiceRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
