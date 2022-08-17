package handlers

import (
	"context"
	"fmt"
	saga "github.com/dislinked/common/saga/messaging"
	events "github.com/dislinked/common/saga/update_order"
	"github.com/dislinked/job_service/application"
	"github.com/dislinked/job_service/domain"
)

type UpdateUserCommandHandler struct {
	jobService        *application.JobService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateUserHandler(jobService *application.JobService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateUserCommandHandler, error) {
	o := &UpdateUserCommandHandler{
		jobService:        jobService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err

	}
	return o, nil

}

func (handler *UpdateUserCommandHandler) handle(command *events.UpdateUserCommand) {
	println("Usao je u command handler")

	/*userJobNode := &domain.UserJobNode{
		UserID:   command.User.Id,
		Username: command.User.Username,
		Skills:   command.User.Skills,
	}*/

	print("command type je: ")
	println(command.Type)

	var reply = events.UpdateUserReply{
		User:    command.User,
		UserOld: command.OldUser,
	}
	switch command.Type {
	case events.UpdateJobNode:
		println("Updating jobnode")
		err, _ := handler.jobService.UpdateUser(context.TODO(), *mapEventUpdateUserToDomainUser(command.User))
		if err.Status != 200 {
			reply.Type = events.JobNodeFailedToUpdate
			println("Failed to updtae update_user_handler " + string(err.Status))
		} else {
			reply.Type = events.JobNodeUpdated
			println("user node updated")
		}
	case events.RollbackJobNode:
		println("Rollback job node")
		err, _ := handler.jobService.UpdateUser(context.TODO(), *mapEventUpdateUserToDomainUser(command.OldUser))
		if err.Status != 200 {
			reply.Type = events.JobNodeRolledBack
			println("Job node rolled back")
		}

	default:
		println("Unknown reply je u pitanju")
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}

	fmt.Println("dosao do kraja user command handler-a job servisa")

}

func mapEventUpdateUserToDomainUser(user events.UserDetails) *domain.UserJobNode {
	domainUser := &domain.UserJobNode{
		UserID:   user.Id,
		Username: user.Username,
	}

	for _, skill := range user.Skills {
		domainUser.Skills = append(domainUser.Skills, skill)
	}

	return domainUser

}
