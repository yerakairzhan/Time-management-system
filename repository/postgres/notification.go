package postgres

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"context"
	_ "database/sql"
)

type NotificationRepository struct {
	q *sqlc.Queries
}

func NewNotificationRepository(q *sqlc.Queries) *NotificationRepository {
	return &NotificationRepository{q: q}
}

func (r *NotificationRepository) Create(notification sqlc.Notification) (int, error) {
	ctx := context.Background()

	arg := sqlc.CreateNotificationParams{
		UserID:  notification.UserID,
		Title:   notification.Title,
		Message: notification.Message,
	}
	id, err := r.q.CreateNotification(ctx, arg)
	return int(id.Int32), err
}
