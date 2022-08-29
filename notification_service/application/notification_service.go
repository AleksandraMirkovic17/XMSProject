package application

import (
	"NotificationService/domain"
)

type NotificationService struct {
	store domain.NotificationStore
}

func NewNotificationService(store domain.NotificationStore) *NotificationService {
	return &NotificationService{
		store: store,
	}
}

func (service *NotificationService) Insert(notification *domain.Notification) error {
	service.store.Insert(notification)
	return nil
}
