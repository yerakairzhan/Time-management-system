-- name: CreateTask :one
INSERT INTO tasks (user_id, name, description, category, priority, deadline)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;