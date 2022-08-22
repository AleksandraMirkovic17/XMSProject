package api

import (
	"MessageService/domain"
	"time"

	pb "github.com/dislinked/common/proto/message_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapMessageFromDomainToPb(message *domain.Message) *pb.UserMessage {
	messagePb := &pb.UserMessage{
		Id:       message.Id.Hex(),
		FromUser: message.FromUser,
		ToUser:   message.ToUser,
		Content:  message.Content,
		Date:     timestamppb.New(message.Date),
	}
	return messagePb
}

func mapMessageFromPbToDomain(messagePb *pb.UserMessage) *domain.Message {
	messageDomain := &domain.Message{
		Id:       primitive.NewObjectID(),
		FromUser: messagePb.FromUser,
		ToUser:   messagePb.ToUser,
		Content:  messagePb.Content,
		Date:     time.Now(),
	}
	return messageDomain
}
