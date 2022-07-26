package api

import (
	"ConnectionService/application"
	"fmt"
	events "github.com/dislinked/common/saga/create_order"
	saga "github.com/dislinked/common/saga/messaging"
)

type RegisterUserCommandHandler struct {
	connectionService *application.ConnectionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(connectionService *application.ConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserCommandHandler, error) {
	o := &RegisterUserCommandHandler{
		connectionService: connectionService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *RegisterUserCommandHandler) handle(command *events.RegisterConnectionUserCommand) {

	fmt.Println("usao u user command handler connection servisa")
	//ovde bi trebala konverzija jednog usera u drugi user

	reply := events.RegisterUserConnectionReply{User: command.User}
	println("Id je" + command.User.Id + " , a public je: ")
	println("Command type je" + string(command.Type))
	switch command.Type {

	case events.UserProfileCreate:
		println("event u register_user_handler je UserProfileCreate")
		err, _ := handler.connectionService.Register(command.User.Id, command.User.Username, command.User.IsPublic)
		if err.Status != 201 {
			reply.Type = events.UserNodeFailedToCreate
			println("Failed to create register_user_handler " + string(err.Status))
		} else {
			reply.Type = events.UserNodeCreated
			println("created register_user_handler ")
		}
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}

	fmt.Println("dosao do kraja user command handler-a connection servisa")

}
