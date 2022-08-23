package orchestrators

import (
	"UserService/domain"
	"UserService/infrastructure/mappers"
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

func (o *UserOrchestrator) nextCommandType(reply events.RegisterUserReplyType) events.RegisterUserCommandType {
	switch reply {
	case events.UserProfileCreated:
		println("Kreiran je user profile pa idemo na auth")
		return events.AuthenticationServiceRegisterUpdate
	case events.UserProfileNotCreated:
		return events.CancelRegistration
	case events.UserProfileRegisterRolledBack:
		return events.CancelRegistration

	case events.AuthenticationServiceUserCreated:
		println("authentication service je updatovan, treba da se kreira user node")
		return events.CreateUserNode
	case events.AuthenticationServiceUserNotCreated:
		return events.RollebackRegisterUserProfile
	case events.AuthenticationServiceRegisterRolledBack:
		return events.RollebackRegisterUserProfile

	case events.UserNodeCreated:
		println("user node kreiran treba sada job node da se kreira")
		return events.CreateJobNode
	case events.UserNodeFailedToCreate:
		println("User node u connections nije kreiran idemo na rollback auth servis!")
		return events.RollbackRegisterAuthenticationService
	case events.ConnectionsRolledBack:
		println("Servis konekcije je rollback idemo na rollback suthentication servids")
		return events.RollbackRegisterAuthenticationService

	case events.JobNodeCreated:
		println("job node je kreiran")
		return events.ApproveRegistration
	case events.JobNodeFailedToCreate:
		println("Job node nije kreiran, rollback connection noda!")
		return events.RollebackRegisterConnectionNode
	case events.RegistrationApproved:
		return events.UnknownCommand
	case events.RegistrationCancelled:
		return events.UnknownCommand

	default:
		return events.UnknownCommand
	}
}

func (o *UserOrchestrator) CreateUser(user *domain.User) error {
	events := &events.RegisterUserCommand{
		Type: events.UserProfileCreate,
		User: *(mappers.MapDomainUserToCommandUser(user)),
	}
	return o.commandPublisher.Publish(events)
}
