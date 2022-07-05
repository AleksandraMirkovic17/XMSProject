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
	id, err := primitive.ObjectIDFromHex(command.Order.Id)
	if err != nil {
		return
	}
	order := &domain.User{ID: id}

	reply := events.RegisterUserReply{Order: command.Order}

	switch command.Type {

	case events.RollbackCreateUserProfile:
		fmt.Println("Auth service:Rollback kredencijala")
		err := handler.orderService.DeleteById(order.ID)
		if err != nil {
			return
		}
		reply.Type = events.UserNotRegistered
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
