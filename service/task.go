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
	return s.taskRepo.Create(userId, task)
}

func (s *TaskServiceImpl) GetTaskByID(userId, taskId int) (sqlc.Task, error) {
	return s.taskRepo.GetByID(userId, taskId)
}

func (s *TaskServiceImpl) GetAllTasks(userId int) ([]sqlc.Task, error) {
	return s.taskRepo.GetAll(userId)
}

func (s *TaskServiceImpl) UpdateTask(userId, taskId int, task sqlc.Task) error {
	return s.taskRepo.Update(userId, taskId, task)
}

func (s *TaskServiceImpl) DeleteTask(userId, taskId int) error {
	return s.taskRepo.Delete(userId, taskId)
}
