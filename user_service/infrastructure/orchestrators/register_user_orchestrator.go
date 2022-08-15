package orchestrators

import (
	"UserService/domain"
	"UserService/infrastructure/mappers"
	"fmt"
	events "github.com/dislinked/common/saga/create_order"
	saga "github.com/dislinked/common/saga/messaging"
)

type UserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*UserOrchestrator, error) {
	o := &UserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle) //slusa odgovore
	println("slusa dgovore ")
	if err != nil {
		return nil, err
	}
	err = o.replySubscriber.Subscribe(o.handleConnection) //slusa odgovore connection servisa nadam se
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *UserOrchestrator) handle(reply *events.RegisterUserReply) {
	println("u orkestratoru u user servisu, metoda handle")
	command := events.RegisterUserCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		println("Publisovanje komande u orkestratoru u user servisu metoda handle")
		_ = o.commandPublisher.Publish(command)
	}
}

/*func (o *UserOrchestrator) handle(reply *events.RegisterUserReply) {
	command := events.RegisterUserCommand{UserAuthentication: reply.UserAuthentication}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}*/

func (o *UserOrchestrator) nextCommandType(reply events.RegisterUserReplyType) events.RegisterUserCommandType {
	switch reply {
	case events.UserProfileCreated:
		println("Kreiran je user profile pa idemo na auth")
		return events.AuthenticationServiceUpdate
	case events.UserProfileNotCreated:
		return events.CancelRegistration

	case events.AuthenticationServiceUpdated:
		println("authentication service je updatovan, treba da se kreira user node")
		return events.CreateUserNode
	case events.AuthenticationServiceNotUpdated:
		return events.RollebackUserProfile
	case events.AuthenticationServiceRolledBack:
		return events.RollebackUserProfile

	case events.UserNodeCreated:
		println("user node kreiran treba sada job node da se kreira")
		return events.CreateJobNode
	case events.UserNodeFailedToCreate:
		return events.RollbackAuthenticationService
	case events.UserProfileRolledBack:
		return events.RollbackAuthenticationService

	case events.JobNodeCreated:
		println("job node je kreiran")
		return events.ApproveRegistration
	case events.JobNodeFailedToCreate:
		return events.RollebackConnectionNode

	default:
		return events.UnknownCommand
	}
}

func (o *UserOrchestrator) handleConnection(reply *events.RegisterUserConnectionReply) events.RegisterUserReplyType {
	//TODO:We check what is the next command type
	//TODO:handle rollback if needed
	fmt.Println("Reply type is", reply.Type)
	switch reply.Type {
	case events.UserProfileCreated:

	}
	return events.UnknownReply

}

func (o *UserOrchestrator) CreateUser(user *domain.User) error {
	events := &events.RegisterUserCommand{
		Type: events.UserProfileCreate,
		User: *(mappers.MapDomainUserToCommandUser(user)),
	}
	return o.commandPublisher.Publish(events)
}

func (o *UserOrchestrator) CreateConnectionUser(user *domain.User) error {
	println("Kreiranje connection usera")
	event := &events.RegisterConnectionUserCommand{
		User: *(mappers.MapDomainUserToConnectionCommandUser(user)),
		Type: events.CreateUserNode,
	}
	println("Pre publishovanja")
	return o.commandPublisher.Publish(event)
}

func (o *UserOrchestrator) CreateJobUser(user *domain.User) error {
	println("Kreiranje job usera")
	event := &events.RegisterJobUserCommand{
		User: *(mappers.MapDomainUserToJobCommandUser(user)),
		Type: events.CreateJobNode,
	}

	println("Pre publishovanja eventa za kreiranje job user noda")
	return o.commandPublisher.Publish(event)

}
