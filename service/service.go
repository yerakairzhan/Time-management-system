package service

import (
	db "TimeManagementSystem/db/sqlc"
)

type Authorization interface {
	CreateUser(user db.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUserByEmail(email string) (db.User, error)
}

type TaskService interface {
	CreateTask(userId int, task db.Task) (int, error)
	GetTasksByUserID(userId int) ([]db.Task, error)
	GetTaskById(id int) (db.Task, error)
	UpdateTask(taskId int, task db.Task) error
	DeleteTask(taskId int) error
}
