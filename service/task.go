package service

import (
	sqlc "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/repository"
)

type TaskServiceImpl struct {
	taskRepo repository.TaskRepository
	userRepo repository.UserRepository
}

func NewTaskService(taskRepo repository.TaskRepository, userRepo repository.UserRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

func (s *TaskServiceImpl) CreateTask(userId int, task sqlc.Task) (int, error) {
	return s.taskRepo.Create((userId), task)
}

func (s *TaskServiceImpl) GetTasksByUserID(userId int) ([]sqlc.Task, error) {
	return s.taskRepo.GetTasksByUserID(userId)
}

func (s *TaskServiceImpl) GetTaskById(taskID int) (sqlc.Task, error) {
	return s.taskRepo.GetTaskById(taskID)
}

func (s *TaskServiceImpl) UpdateTask(taskId int, task sqlc.Task) error {
	return s.taskRepo.Update(taskId, task)
}

func (s *TaskServiceImpl) DeleteTask(taskId int) error {
	return s.taskRepo.Delete(taskId)
}
