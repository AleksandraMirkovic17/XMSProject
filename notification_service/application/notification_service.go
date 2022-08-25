package application

import (
	"NotificationService/domain"
	orchestrators "NotificationService/infrastructure/orchestrator"
)

type NotificationService struct {
	store        domain.NotificationStore
	orchestrator *orchestrators.CreateNotificationOrchestrator
}

func NewNotificationService(store domain.NotificationStore, orchestrator *orchestrators.CreateNotificationOrchestrator) *NotificationService {
	return &NotificationService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *NotificationService) Insert(notification *domain.Notification) error {
	service.store.Insert(notification)
	return nil
}
