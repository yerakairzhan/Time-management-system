package service

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/repository"
)

type NotificationServiceImpl struct {
	repo repository.NotificationRepository
}

func NewNotificationService(repo repository.NotificationRepository) NotificationService {
	return &NotificationServiceImpl{repo: repo}
}

func (n *NotificationServiceImpl) CreateNotification(notification sqlc.Notification) (int, error) {
	return n.repo.Create(notification)
}
