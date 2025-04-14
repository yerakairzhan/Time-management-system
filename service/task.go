package service

import (
	"context"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository, userRepo repository.UserRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, name string) error {
	// Логика создания задачи
	return s.repo.Create(ctx, name)
}

func (s *TaskService) ListTasks(ctx context.Context) ([]Task, error) {
	// Логика получения списка задач
	return s.repo.GetAll(ctx)
}
