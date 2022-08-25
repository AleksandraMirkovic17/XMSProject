package orchestrator

import (
	events "github.com/dislinked/common/saga/create_notification"
	saga "github.com/dislinked/common/saga/messaging"
)

type CreateNotificationOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewCreateNotificationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*CreateNotificationOrchestrator, error) {
	o := &CreateNotificationOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *CreateNotificationOrchestrator) handle(reply *events.CreateNotificationReply) {
	command := events.CreateNotificationCommand{Notification: reply.Notification}
	//command.Type = o.nextCommandType(reply.Type)
	command.Type = events.CreateNotification
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *CreateNotificationOrchestrator) Start(NotificationDetails events.NotificationDetails) error {
	event := &events.CreateNotificationCommand{
		Type:         events.CreateNotification,
		Notification: NotificationDetails,
	}

	return o.commandPublisher.Publish(event)
}

/*
func (o *RegisterUserOrchestrator) nextCommandType(reply events.RegisterUserReplyType) events.RegisterUserCommandType {
	switch reply {
	case events.UserProfileCreated:
		println("Kreiran je user profile pa idemo na auth")
		return events.AuthenticationServiceRegisterUpdate
	case events.UserProfileNotCreated:
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
		return events.RollbackRegisterAuthenticationService
	case events.UserProfileRegisterRolledBack:
		return events.RollbackRegisterAuthenticationService

	case events.JobNodeCreated:
		println("job node je kreiran")
		return events.ApproveRegistration
	case events.JobNodeFailedToCreate:
		return events.RollebackRegisterConnectionNode

	default:
		return events.UnknownCommand
	}
}*/
