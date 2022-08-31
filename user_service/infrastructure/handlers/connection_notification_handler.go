package handlers

import (
	"UserService/application"
	events "github.com/dislinked/common/saga/connection_notification"
	saga "github.com/dislinked/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectionNotificationHandler struct {
	service           *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewConnectionNotificatioHandler(service *application.UserService, publiser saga.Publisher, subscriber saga.Subscriber) (*ConnectionNotificationHandler, error) {
	o := &ConnectionNotificationHandler{
		service:           service,
		replyPublisher:    publiser,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *ConnectionNotificationHandler) handle(command events.ConnectionNotificationCommand) {
	println("Nalazim se u hendleru user servisa za slanje notifikacija za konekcije")
	print("Command type je: ")
	println(command.Type)
	reply := &events.ConnectionNotificationReply{
		Notification: command.Notification,
	}

	switch command.Type {
	case events.CheckAndGetReceiverById:
		println("Checking notification settings")
		//TO DO: dodati
		senderObjectId, err := primitive.ObjectIDFromHex(command.Notification.SenderId)
		if err != nil {
			println("Hex to object sender id converting error!")
			reply.Type = events.NotificationsAreTurnedOff
			break
		}
		receiverObjectId, err := primitive.ObjectIDFromHex(command.Notification.ReceiverId)
		if err != nil {
			println("Hex to object receiver id converting error!")
			reply.Type = events.NotificationsAreTurnedOff
			break
		}
		senderUser, err1 := handler.service.GetOne(senderObjectId)
		receiverUser, err2 := handler.service.GetOne(receiverObjectId)
		if err1 != nil || err2 != nil {
			reply.Type = events.NotificationsAreTurnedOff
			break
		}
		if !receiverUser.NotificationsFollowedProfiles {
			reply.Type = events.NotificationsAreTurnedOff
			break
		}
		reply.Notification.Sender = senderUser.Name + " " + senderUser.Surname + "(@" + senderUser.Username + ")"
		reply.Notification.Receiver = receiverUser.Name + " " + receiverUser.Surname + "(@" + receiverUser.Username + ")"
		if command.Notification.Request {
			reply.Notification.Content = receiverUser.Name + ", " + senderUser.Name + " " + senderUser.Surname + "(@" + senderUser.Username + ") sent you a friend request!"
		} else {
			reply.Notification.Content = receiverUser.Name + ", " + senderUser.Name + " " + senderUser.Surname + "(@" + senderUser.Username + ") connected with you!"
		}
		reply.Type = events.ReceiverFilled
		break
	default:
		reply.Type = events.UnknownReply
		break

	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
