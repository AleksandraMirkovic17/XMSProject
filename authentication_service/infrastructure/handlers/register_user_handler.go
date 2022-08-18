package handlers

import (
	"AuthenticationService/application"
	"AuthenticationService/domain"
	"AuthenticationService/infrastructure/mappers"
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
	println("Nalazim se u handleru authentication servisa")
	id, err := primitive.ObjectIDFromHex(command.User.Id)
	if err != nil {
		return
	}
	order := &domain.UserAuthentication{ID: id} //ovde se postavlja id

	reply := events.RegisterUserReply{User: command.User}
	print("Pre switchovanja")
	print(command.Type)
	switch command.Type {
	case events.AuthenticationServiceRegisterUpdate:
		println("Update authentication servisa")
		_, err := handler.orderService.Register(mappers.MapCommandToAuthUser(command))
		if err != nil {
			println("Nije updatovan")
			reply.Type = events.AuthenticationServiceUserNotCreated
			return
		}
		println("uspeso update authentication servis")
		reply.Type = events.AuthenticationServiceUserCreated

	case events.RollbackRegisterAuthenticationService:
		fmt.Println("Auth service:Rollback authentication servisa")
		err := handler.orderService.DeleteById(order.ID)
		if err != nil {
			return
		}
		reply.Type = events.AuthenticationServiceRegisterRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
