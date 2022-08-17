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
		User: reply.User,
		Type: o.nextCommand(reply.Type),
	}
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}

}

func (o *UpdateUserOrchestrator) nextCommand(reply events.UpdateUserReplyType) events.UpdateUserCommandType {
	switch reply {
	case events.UserProfileUpdated:
		return events.AuthenticationServiceUpdate
	case events.UserProfileNotUpdated:
		return events.CancelUpdate

	case events.AuthenticationServiceUpdated:
		return events.UpdateUserNode
	case events.AuthenticationServiceNotUpdated:
		return events.RollebackUserProfile
	case events.AuthenticationServiceRolledBack:
		return events.RollebackUserProfile

	case events.UserNodeUpdated:
		return events.UpdateJobNode
	case events.UserNodeFailedToUpdate:
		return events.RollbackAuthenticationService
	case events.ConnectionsRolledBack:
		return events.RollbackAuthenticationService

	case events.JobNodeUpdated:
		return events.ApproveUpdate
	case events.JobNodeFailedToUpdate:
		return events.RollebackConnectionNode
	default:
		return events.UnknownCommand

	}

}

func (o *UpdateUserOrchestrator) Start(user *events.UserDetails, oldUser *events.UserDetails) error {
	event := &events.UpdateUserCommand{
		User:    *user,
		OldUser: *oldUser,
		Type:    events.AuthenticationServiceUpdate,
	}

	return o.commandPublisher.Publish(event)

}
