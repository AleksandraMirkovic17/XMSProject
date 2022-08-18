package orchestrators

import (
	saga "github.com/dislinked/common/saga/messaging"
	events "github.com/dislinked/common/saga/update_order"
)

type UpdateUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewUpdateUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateUserOrchestrator, error) {
	o := &UpdateUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *UpdateUserOrchestrator) handle(reply *events.UpdateUserReply) {
	println("Nalazim se u hendleru update orkestratora")
	command := events.UpdateUserCommand{
		User:    reply.User,
		OldUser: reply.UserOld,
		Type:    o.nextCommand(reply.Type),
	}
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
		println("Publisovana je komanda")
	}

}

func (o *UpdateUserOrchestrator) nextCommand(reply events.UpdateUserReplyType) events.UpdateUserCommandType {
	switch reply {
	case events.UserProfileUpdated:
		println("User Update idemo na auth update")
		return events.AuthenticationServiceUpdate
	case events.UserProfileNotUpdated:
		return events.CancelUpdate

	case events.AuthenticationServiceUpdated:
		println("Authentication service je updated pa idemo na pdate user noda")
		return events.UpdateUserNode
	case events.AuthenticationServiceNotUpdated:
		println("Authentication nije udated pa moramo na rollback user profila onda")
		return events.RollebackUserProfile
	case events.AuthenticationServiceRolledBack:
		return events.RollebackUserProfile

	case events.UserNodeUpdated:
		return events.UpdateJobNode
	case events.UserNodeFailedToUpdate:
		return events.RollbackAuthenticationService
	case events.ConnectionsRolledBack:
		println("Rollback connection servisa je odradjen idemo na rollback authentication servisa")
		return events.RollbackAuthenticationService

	case events.JobNodeUpdated:
		return events.ApproveUpdate
	case events.JobNodeFailedToUpdate:
		println("Job node se nije update radimo rollback connection servisa")
		return events.RollebackConnectionNode
	default:
		return events.UnknownCommand

	}

}

func (o *UpdateUserOrchestrator) Start(user *events.UserDetails, oldUser *events.UserDetails) error {

	println("startovanje orkestrotora")
	event := &events.UpdateUserCommand{
		User:    *user,
		OldUser: *oldUser,
		Type:    events.AuthenticationServiceUpdate,
	}

	return o.commandPublisher.Publish(event)

}
