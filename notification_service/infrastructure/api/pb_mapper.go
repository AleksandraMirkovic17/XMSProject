package api

import (
	"NotificationService/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/dislinked/common/proto/notification_service"
)

func mapNotificationToPB(notification *domain.Notification) *pb.Notification {
	notificationPb := &pb.Notification{
		Id:      notification.Id.Hex(),
		User:    notification.User,
		Content: notification.Content,
		Url:     notification.Url,
		Seen:    notification.Seen,
		Date:    timestamppb.New(notification.Date),
	}
	return notificationPb
}

func mapNotificationToDomain(notification *pb.Notification) *domain.Notification {
	id, err := primitive.ObjectIDFromHex((*notification).Id)
	if err != nil {
		return nil
	}
	userDomain := &domain.Notification{
		Id:      id,
		User:    notification.User,
		Content: notification.Content,
		Url:     notification.Url,
		Seen:    notification.Seen,
		Date:    time.Now(),
	}
	return userDomain
}

/*func mapNotificationCommandToDomain(notification *events.NotificationDetails) *domain.Notification {
	id, err := primitive.ObjectIDFromHex((*notification).Id)
	if err != nil {
		return nil
	}
	userDomain := &domain.Notification{
		Id:      id,
		User:    notification.User,
		Content: notification.Content,
		Url:     notification.Url,
		Seen:    notification.Seen,
		Date:    time.Now(),
	}
	return userDomain
}*/
