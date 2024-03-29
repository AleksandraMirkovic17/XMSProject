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

func (handler *RegisterUserCommandHandler) handle(command *events.RegisterUserCommand) {

	fmt.Println("usao u user command handler connection servisa")
	//ovde bi trebala konverzija jednog usera u drugi user

	reply := events.RegisterUserReply{User: command.User}
	println("Username: "+command.User.Username+", id: ", command.User.Id+", public:")
	println("Command type je" + string(command.Type))
	println(command.Type)
	switch command.Type {

	case events.CreateUserNode:
		println("event u register_user_handler je UserProfileCreate")
		err, _ := handler.connectionService.Register(command.User.Id, command.User.Username, command.User.IsPublic)
		if err.Status != 201 {
			reply.Type = events.UserNodeFailedToCreate
			print(err.Status)
			println(" Failed to create register_user_handler " + string(err.Status))
			break
		} else {
			reply.Type = events.UserNodeCreated
			println("created register_user_handler ")
			break
		}
	case events.RollebackRegisterConnectionNode:
		println("Nalazim se u rollback connection node, medjutim nije implamentirana metoda")
		reply.Type = events.ConnectionsRolledBack
		break

	default:
		println("Unknown reply je u pitanju")
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}

	fmt.Println("dosao do kraja user command handler-a connection servisa")

}
