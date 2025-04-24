-- name: CreateNotification :one
INSERT INTO notifications (user_id, title, message)
VALUES ($1, $2, $3)
    RETURNING user_id;

-- name: GetNotificationByUserId :many
select * from notifications where user_id = $1;

-- name: GetNotificationById :one
select * from notifications where id = $1;

-- name: UpdateNotification :exec
update notifications set title = $1, message = $2 where id = $3;

-- name: DeleteNotification :exec
delete from notifications where id = $1;