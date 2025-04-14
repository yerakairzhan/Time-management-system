package postgres

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"context"
	"time"
)

type TaskRepository struct {
	q *sqlc.Queries
}

func NewTaskRepository(q *sqlc.Queries) *TaskRepository {
	return &TaskRepository{q: q}
}

func (r *TaskRepository) CreateTask(ctx context.Context, name, description, category, priority string, deadline time.Time) (sqlc.Task, error) {
	arg := sqlc.CreateTaskParams{
		Name:        name,
		Description: description,
		Category:    category,
		Priority:    priority,
		Deadline:    deadline,
	}
	return r.q.CreateTask(ctx, arg)
}

func (r *TaskRepository) ListTasks(ctx context.Context) ([]sqlc.Task, error) {
	return r.q.ListTasks(ctx)
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id int32) (sqlc.Task, error) {
	return r.q.GetTaskByID(ctx, id)
}
