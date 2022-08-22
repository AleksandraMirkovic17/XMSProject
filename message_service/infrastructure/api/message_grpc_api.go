package api

import (
	"MessageService/application"
	"context"

	pb "github.com/dislinked/common/proto/message_service"
	"google.golang.org/grpc/status"
)

type MessageHandler struct {
	pb.UnimplementedMessageServiceServer
	service *application.MessageService
}

func NewMessageHandler(service *application.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

func (handler *MessageHandler) SendMessage(ctx context.Context, request *pb.NewUserMessage) (*pb.NewUserMessage, error) {
	message := mapMessageFromPbToDomain(request.UserMessage)
	newMessage, err := handler.service.Insert(message)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	response := &pb.NewUserMessage{
		UserMessage: mapMessageFromDomainToPb(newMessage),
	}

	return response, nil
}

func (handler *MessageHandler) GetByUser(ctx context.Context, request *pb.GetMessagesByUserRequest) (*pb.GetMultipleMessages, error) {
	uuidUser := request.Uuid
	messages, err := handler.service.GetAllByUser(uuidUser)
	if err != nil {
		return nil, err
	}
	response := &pb.GetMultipleMessages{
		UserMessages: []*pb.UserMessage{},
	}
	for _, message := range messages {
		current := mapMessageFromDomainToPb(message)
		response.UserMessages = append(response.UserMessages, current)
	}
	return response, nil
}
