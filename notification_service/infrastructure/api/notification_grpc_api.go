package api

import (
	"NotificationService/application"
	"NotificationService/infrastructure/orchestrator"
	"context"

	pb "github.com/dislinked/common/proto/notification_service"
	events "github.com/dislinked/common/saga/create_notification"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationServiceServer
	service                        *application.NotificationService
	CreateNotificationOrchestrator *orchestrator.CreateNotificationOrchestrator
}

func NewNotificationHandler(service *application.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

func (handler *NotificationHandler) CreateNotification(ctx context.Context, request *pb.NewNotification) (*pb.NewNotification, error) {
	notification := mapNotificationToDomain(request.Notification)
	handler.service.Insert(notification)
	handler.CreateNotificationOrchestrator.Start(events.NotificationDetails{
		Id:      request.Notification.Id,
		User:    request.Notification.User,
		Content: request.Notification.Content,
		Url:     request.Notification.Url,
		Seen:    false,
		//Date:    time.Now(),
	})

	/*response := &pb.NewNotification{
		Notification.Content: "TEST",
	}*/
	return nil, nil
}

func (handler *NotificationHandler) GetByUser(ctx context.Context, request *pb.GetNotificationsByUserRequest) (*pb.GetMultipleNotifications, error) {
	uuidUser := request.Uuid
	notifications, err := handler.service.GetAllByUser(uuidUser)
	if err != nil {
		return nil, err
	}
	response := &pb.GetMultipleNotifications{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotificationToPB(notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}
