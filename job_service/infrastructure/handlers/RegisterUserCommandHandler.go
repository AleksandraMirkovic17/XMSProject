package handlers

import (
	"fmt"
	events "github.com/dislinked/common/saga/create_order"
	saga "github.com/dislinked/common/saga/messaging"
	"github.com/dislinked/job_service/application"
	"golang.org/x/net/context"
)

type RegisterUserCommandHandler struct {
	jobService        *application.JobService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewJobHandler(jobService *application.JobService, publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserCommandHandler, error) {
	o := &RegisterUserCommandHandler{
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

func (handler *RegisterUserCommandHandler) handle(command *events.RegisterUserCommand) {
	println("Usao je u command handler")

	/*userJobNode := &domain.UserJobNode{
		UserID:   command.User.Id,
		Username: command.User.Username,
		Skills:   command.User.Skills,
	}*/

	print("command type je: ")
	println(command.Type)

	var reply = events.RegisterUserReply{}
	switch command.Type {
	case events.CreateJobNode:
		println("Creta Job Node")
		err, _ := handler.jobService.CreateUser(context.TODO(), command.User.Id, command.User.Username)
		if err.Status != 201 {
			reply.Type = events.JobNodeFailedToCreate
			println("Failed to create register_user_handler " + string(err.Status))
		} else {
			reply.Type = events.JobNodeCreated
			println("user node created")
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
