-- name: CreateNotification :one
INSERT INTO notifications (user_id, title, message)
VALUES ($1, $2, $3)
    RETURNING user_id;