package handlers

import (
	"UserService/application"
	"UserService/infrastructure/mappers"
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

//treba da se nalazi u connection servisu
//treba orkestrator umesto ovoga
func (handler *CreateOrderCommandHandler) handle(command *events.RegisterUserCommand) {
	println("Nalazim se u hendleru user servisa za creatw")
	order := mappers.MapCommandToUser(command) //ovde se postavlja id

	reply := events.RegisterUserReply{User: command.User}

	switch command.Type {
	case events.UserProfileCreate:
		println("komanda jr userProfileCreate")
		err := handler.userService.Register(order)
		if err != nil {
			reply.Type = events.UserProfileNotCreated
			return
		}
		reply.Type = events.UserProfileCreated

	case events.RollebackRegisterUserProfile:
		fmt.Println("Auth service:Rollback user servisa")
		id, err := primitive.ObjectIDFromHex(command.User.Id)
		if err != nil {
			println("Desila se greska prilikom konvertovanja brisanja postojeces")
			return
		}
		toDelete, _ := handler.userService.GetOne(id)
		handler.userService.Delete(toDelete)
		reply.Type = events.UserProfileRegisterRolledBack
	case events.ApproveRegistration:
		fmt.Println("Approve registration")
		reply.Type = events.RegistrationApproved
	case events.CancelRegistration:
		reply.Type = events.RegistrationCancelled

	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
		println("Reply je publisovan")
	}
}
