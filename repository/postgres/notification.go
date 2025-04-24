package postgres

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"context"
	"database/sql"
	_ "database/sql"
	"log"
)

type NotificationRepository struct {
	q *sqlc.Queries
}

func (r *NotificationRepository) GetNotificationById(id int) (sqlc.Notification, error) {
	ctx := context.Background()
	return r.q.GetNotificationById(ctx, int32(id))
}

func (r *NotificationRepository) Update(n sqlc.Notification) error {
	ctx := context.Background()
	arg := sqlc.UpdateNotificationParams{
		Title:   n.Title,
		Message: n.Message,
		ID:      n.ID,
	}

	log.Println(n)

	return r.q.UpdateNotification(ctx, arg)
}

func (r *NotificationRepository) Delete(id int) error {
	ctx := context.Background()
	return r.q.DeleteNotification(ctx, int32(id))
}

func (r *NotificationRepository) GetNotificationsByUserID(userId int) ([]sqlc.Notification, error) {
	ctx := context.Background()
	return r.q.GetNotificationByUserId(ctx, sql.NullInt32{Int32: int32(userId), Valid: true})
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
