package postgres

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"context"
	"database/sql"
)

type TaskRepository struct {
	q *sqlc.Queries
}

func (r *TaskRepository) Create(userId int, task sqlc.Task) (int, error) {
	ctx := context.Background()

	arg := sqlc.CreateTaskParams{
		UserID:      int32(userId),
		Name:        task.Name,
		Description: task.Description,
		Category:    task.Category,
		Priority:    task.Priority,
		Deadline:    task.Deadline,
	}

	ans, err := r.q.CreateTask(ctx, arg)
	return int(ans), err
}

func (r *TaskRepository) GetTaskById(TaskID int) (sqlc.Task, error) {
	ctx := context.Background()
	return r.q.GetTaskByID(ctx, int32(TaskID))
}

func (r *TaskRepository) Update(taskId int, task sqlc.Task) error {
	ctx := context.Background()
	arg := sqlc.UpdateTaskParams{
		ID:          int32(taskId),
		Name:        sql.NullString{String: task.Name, Valid: true},
		Description: task.Description,
		Category:    sql.NullString{String: task.Category, Valid: true},
		Priority:    sql.NullString{String: task.Priority, Valid: true},
		UserID:      task.UserID,
		Deadline:    task.Deadline,
	}
	return r.q.UpdateTask(ctx, arg)
}

func (r *TaskRepository) Delete(taskId int) error {
	return r.q.DeleteTask(context.Background(), int32(taskId))
}

func NewTaskRepository(q *sqlc.Queries) *TaskRepository {
	return &TaskRepository{q: q}
}

func (r *TaskRepository) ListTasks(ctx context.Context) ([]sqlc.Task, error) {
	return r.q.ListTasks(ctx)
}

func (r *TaskRepository) GetTasksByUserID(id int) ([]sqlc.Task, error) {
	ctx := context.Background()
	return r.q.GetTasksByUserID(ctx, int32(id))
}
