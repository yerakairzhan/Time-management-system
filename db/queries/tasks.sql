-- name: CreateTask :one
INSERT INTO tasks (user_id, name, description, category, priority, deadline)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING user_id;

-- name: ListTasks :many
SELECT *
FROM tasks;

-- name: GetTasksByUserID :many
SELECT *
FROM tasks
WHERE user_id = $1;

-- name: GetTaskByID :one
SELECT id, user_id, name, description, category, priority, deadline
FROM tasks
WHERE id = $1;

-- name: UpdateTask :exec
UPDATE tasks
SET
    name = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(sqlc.narg('description'), description),
    category = COALESCE(sqlc.narg('category'), category),
    priority = COALESCE(sqlc.narg('priority'), priority),
    deadline = COALESCE(sqlc.narg('deadline'), deadline)
WHERE id = sqlc.arg('id') AND user_id = sqlc.arg('user_id');

-- name: DeleteTask :exec
delete from tasks where id = $1;
