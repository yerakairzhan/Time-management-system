package service

import (
	db "TimeManagementSystem/db/sqlc"
)

type Authorization interface {
	CreateUser(user db.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TaskService interface {
	CreateTask(userId int, task db.Task) (int, error)
	GetTaskByID(userId, taskId int) (db.Task, error)
	GetAllTasks(userId int) ([]db.Task, error)
	UpdateTask(userId, taskId int, task db.Task) error
	DeleteTask(userId, taskId int) error
}

type UserRepository interface {
	Create(user db.User) (int, error)
	GetUserByEmail(email string) (db.User, error)
}

type TaskRepository interface {
	Create(userId int, task db.Task) (int, error)
	GetByID(userId, taskId int) (db.Task, error)
	GetAll(userId int) ([]db.Task, error)
	Update(userId, taskId int, task db.Task) error
	Delete(userId, taskId int) error
}

type Service struct {
	Authorization
	TaskService
}

type Repository struct {
	UserRepository
	TaskRepository
}

func NewService(repo Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.UserRepository),
		TaskService:   NewTaskService(repo.TaskRepository, repo.UserRepository),
	}
}
