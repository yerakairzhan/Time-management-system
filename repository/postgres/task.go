package postgres

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"context"
	"database/sql"
	"time"
)

type TaskRepository struct {
	q *sqlc.Queries
}

func (r *TaskRepository) Create(userId int, task sqlc.Task) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TaskRepository) GetByID(userId, taskId int) (sqlc.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TaskRepository) GetAll(userId int) ([]sqlc.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TaskRepository) Update(userId, taskId int, task sqlc.Task) error {
	//TODO implement me
	panic("implement me")
}

func (r *TaskRepository) Delete(userId, taskId int) error {
	//TODO implement me
	panic("implement me")
}

func NewTaskRepository(q *sqlc.Queries) *TaskRepository {
	return &TaskRepository{q: q}
}

func (r *TaskRepository) CreateTask(
	ctx context.Context,
	name, description, category, priority string,
	deadline time.Time,
) (sqlc.Task, error) {
	arg := sqlc.CreateTaskParams{
		Name:        sql.NullString{String: name, Valid: true},
		Description: sql.NullString{String: description, Valid: true},
		Category:    sql.NullString{String: category, Valid: true},
		Priority:    sql.NullString{String: priority, Valid: true},
		Deadline:    sql.NullTime{Time: deadline, Valid: true},
	}

	return r.q.CreateTask(ctx, arg)
}

func (r *TaskRepository) ListTasks(ctx context.Context) ([]sqlc.Task, error) {
	return r.q.ListTasks(ctx)
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id int32) (sqlc.Task, error) {
	return r.q.GetTaskByID(ctx, id)
}
