package domain

type NotificationStore interface {
	Insert(message *Notification) error
	GetAllByUser(uuid string) ([]*Notification, error)
}
