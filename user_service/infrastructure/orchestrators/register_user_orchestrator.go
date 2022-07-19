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
	if err != nil {
		return nil, err
	}
	err = o.replySubscriber.Subscribe(o.handleConnection) //slusa odgovore connection servisa nadam se
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *UserOrchestrator) handle(reply *events.RegisterUserReply) events.RegisterUserReplyType {
	if reply.Type == events.UserProfileRolledBack {
		fmt.Println("User profile rolled back register_user_orchestrator!")
		return events.UserProfileRolledBack
	}
	if reply.Type == events.UserProfileCreated {
		fmt.Println("User profile created in register_user_orchestrator!")
		return events.UserProfileCreated
	}
	/*TODO: mozda treba da se doda i za update*/
	fmt.Println("fkk handle")
	return events.UnknownReply

}

func (o *UserOrchestrator) handleConnection(reply *events.RegisterUserConnectionReply) events.RegisterUserReplyType {
	//TODO:We check what is the next command type
	//TODO:handle rollback if needed
	fmt.Println(reply.Type)
	fmt.Println("BAAAAAAACKS connection")
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
		Type: events.UserProfileCreate,
	}
	println("Pre publishovanja")
	return o.commandPublisher.Publish(event)
}
