package orchestrator

import (
	events "github.com/dislinked/common/saga/create_order"
	saga "github.com/dislinked/common/saga/messaging"
)

type RegisterUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewRegisterUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserOrchestrator, error) {
	o := &RegisterUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *RegisterUserOrchestrator) Start(userDetails events.UserDetails) error {
	event := &events.RegisterUserCommand{
		Type: events.AuthenticationServiceUpdate,
		User: userDetails,
	}

	return o.commandPublisher.Publish(event)
}

func (o *RegisterUserOrchestrator) handle(reply *events.RegisterUserReply) {
	println("U register user servisu orchestraatoru authentication servisa")
	command := events.RegisterUserCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *RegisterUserOrchestrator) nextCommandType(reply events.RegisterUserReplyType) events.RegisterUserCommandType {
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
