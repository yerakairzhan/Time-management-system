package service

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/repository"
)

type NotificationServiceImpl struct {
	repo repository.NotificationRepository
}

func (n *NotificationServiceImpl) GetById(id int) (sqlc.Notification, error) {
	return n.repo.GetNotificationById(id)
}

func (n *NotificationServiceImpl) UpdateNotification(notification sqlc.Notification) error {
	return n.repo.Update(notification)
}

func (n *NotificationServiceImpl) DeleteNotification(notificationId int) error {
	return n.repo.Delete(notificationId)
}

func (n *NotificationServiceImpl) GetNotificationsByUserID(userId int) ([]sqlc.Notification, error) {
	return n.repo.GetNotificationsByUserID(userId)
}

func NewNotificationService(repo repository.NotificationRepository) NotificationService {
	return &NotificationServiceImpl{repo: repo}
}

func (n *NotificationServiceImpl) CreateNotification(notification sqlc.Notification) (int, error) {
	return n.repo.Create(notification)
}
