package repository

import (
	sqlc "TimeManagementSystem/db/sqlc"
)

type TaskRepository interface {
	Create(userId int, task sqlc.Task) (int, error)
	GetTasksByUserID(userId int) ([]sqlc.Task, error)
	GetTaskById(id int) (sqlc.Task, error)
	Update(taskId int, task sqlc.Task) error
	Delete(taskId int) error
}

type Authorization interface {
	CreateUser(user sqlc.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UserRepository interface {
	Create(user sqlc.User) (int, error)
	GetUserByEmail(email string) (sqlc.User, error)
}
