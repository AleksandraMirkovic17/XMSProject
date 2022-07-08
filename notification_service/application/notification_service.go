package application

import (
	"NotificationService/application/adapters"
	"NotificationService/domain"
	"NotificationService/infrastructure/persistence"
	"NotificationService/startup/config"
	"context"
	"fmt"
	"time"

	connectionService "github.com/dislinked/common/proto/connection_service"
	pb "github.com/dislinked/common/proto/notification_service"
	userService "github.com/dislinked/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationService struct {
	store            persistence.NotificationStore
	ConnectionClient connectionService.ConnectionServiceClient
	UserClient       userService.UserServiceClient
}

func NewNotificationService(store persistence.NotificationStore, c *config.Config) *NotificationService {
	return &NotificationService{
		store:            store,
		ConnectionClient: adapters.NewConnectionClient(fmt.Sprintf("%s:%s", c.ConnectionHost, c.ConnectionPort)),
		UserClient:       adapters.NewUserClient(fmt.Sprintf("%s:%s", c.UserHost, c.UserPort)),
	}
}

func (service *NotificationService) GetAllNotifications(ctx context.Context, request *pb.GetAllNotificationsRequest) (*pb.GetAllNotificationsResponse, error) {
	userId := request.UserID
	var userNotifications []*pb.Notification
	allNotifications, err := service.store.GetAll(ctx)

	if err == nil {
		for _, notification := range allNotifications {
			if notification.OwnerId.Hex() == userId {
				var newNotification = pb.Notification{
					OwnerId:      notification.OwnerId.Hex(),
					ForwardUrl:   notification.ForwardUrl,
					Text:         notification.Text,
					Date:         &timestamppb.Timestamp{Seconds: notification.Date.Unix()},
					Seen:         notification.Seen,
					UserFullName: notification.UserFullName,
				}
				userNotifications = append(userNotifications, &newNotification)
			}
		}
	}

	return &pb.GetAllNotificationsResponse{
		Notifications: userNotifications,
	}, err
}

func (service *NotificationService) MarkAllAsSeen(ctx context.Context, request *pb.MarkAllAsSeenRequest) (*pb.MarkAllAsSeenResponse, error) {
	userId := request.UserID
	allNotifications, err := service.store.GetAll(ctx)

	if err == nil {
		for _, notification := range allNotifications {
			if notification.OwnerId.Hex() == userId {
				if !notification.Seen {
					service.store.MarkAsSeen(ctx, notification.Id)
				}
			}
		}
	}

	return &pb.MarkAllAsSeenResponse{UserID: userId}, err
}

func (service *NotificationService) InsertNotification(ctx context.Context, request *pb.InsertNotificationRequest) (*pb.InsertNotificationRequestResponse, error) {
	ownerId, err := primitive.ObjectIDFromHex(request.Notification.OwnerId)

	notification := &domain.Notification{
		OwnerId:      ownerId,
		ForwardUrl:   request.Notification.ForwardUrl,
		Text:         request.Notification.Text,
		Date:         time.Now(),
		Seen:         false,
		UserFullName: request.Notification.UserFullName,
	}

	if service.UserAcceptsNotification(notification) {
		service.store.Insert(ctx, notification)
	}

	return &pb.InsertNotificationRequestResponse{
		Notification: request.Notification,
	}, err
}

func (service *NotificationService) GetUserSettings(ctx context.Context, request *pb.GetUserSettingsRequest) (*pb.GetUserSettingsResponse, error) {
	userId := request.UserID
	id, err := primitive.ObjectIDFromHex(userId)
	if err == nil {
		settings := service.store.GetOrInitUserSetting(ctx, id)
		return &pb.GetUserSettingsResponse{
			UserID:                  settings.OwnerId.Hex(),
			PostNotifications:       settings.PostNotifications,
			ConnectionNotifications: settings.ConnectionNotifications,
			MessageNotifications:    settings.MessageNotifications,
		}, nil
	} else {
		return &pb.GetUserSettingsResponse{}, err
	}
}

func (service *NotificationService) UpdateUserSettings(ctx context.Context, request *pb.UpdateUserSettingsRequest) (*pb.GetUserSettingsResponse, error) {
	userId := request.UserID
	id, err := primitive.ObjectIDFromHex(userId)
	if err == nil {
		settings := domain.UserSettings{
			OwnerId:                 id,
			PostNotifications:       true,
			ConnectionNotifications: true,
			MessageNotifications:    true,
		}

		if request.SettingsCode == "2" {
			settings.PostNotifications = false
			settings.ConnectionNotifications = false
			settings.MessageNotifications = false
		} else if request.SettingsCode == "3" {
			settings.PostNotifications = true
			settings.ConnectionNotifications = false
			settings.MessageNotifications = false
		} else if request.SettingsCode == "4" {
			settings.PostNotifications = false
			settings.ConnectionNotifications = true
			settings.MessageNotifications = false
		} else if request.SettingsCode == "5" {
			settings.PostNotifications = false
			settings.ConnectionNotifications = false
			settings.MessageNotifications = true
		} else if request.SettingsCode == "6" {
			settings.PostNotifications = true
			settings.ConnectionNotifications = false
			settings.MessageNotifications = true
		} else if request.SettingsCode == "7" {
			settings.PostNotifications = false
			settings.ConnectionNotifications = true
			settings.MessageNotifications = true
		} else if request.SettingsCode == "8" {
			settings.PostNotifications = true
			settings.ConnectionNotifications = true
			settings.MessageNotifications = false
		}

		service.store.ModifyOrInsertSetting(ctx, &settings)
		return &pb.GetUserSettingsResponse{
			UserID: settings.OwnerId.Hex(),
		}, nil
	} else {
		return &pb.GetUserSettingsResponse{}, err
	}
}

func (service *NotificationService) UserAcceptsNotification(notification *domain.Notification) bool {
	settings := service.store.GetOrInitUserSetting(context.TODO(), notification.OwnerId)
	if notification.Text == "sent you a message" {
		return settings.MessageNotifications
	} else if notification.Text == "is now your friend" || notification.Text == "sent you a friend request" {
		return settings.ConnectionNotifications
	} else if notification.Text == "posted on their profile" || notification.Text == "commented on your post" {
		return settings.PostNotifications
	}

	return true
}
