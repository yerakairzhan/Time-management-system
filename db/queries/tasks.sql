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
WHERE user_id = $1
ORDER BY priority DESC;

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

---- TIMER AHEAD

-- name: StartTaskTimer :exec
INSERT INTO task_time_logs (task_id, start_time)
VALUES ($1, NOW())
    RETURNING *;

-- name: StopTaskTimer :exec
UPDATE task_time_logs
SET end_time = NOW()
WHERE task_id = $1 AND end_time IS NULL
    RETURNING *;

-- name: GetActiveTimer :one
SELECT * FROM task_time_logs
WHERE task_id = $1 AND end_time IS NULL
    LIMIT 1;

-- name: GetTaskTimeLogs :many
SELECT * FROM task_time_logs
WHERE task_id = $1
ORDER BY start_time DESC;

-- name: GetTimeSpent :many
select start_time , end_time from task_time_logs where task_id = $1;