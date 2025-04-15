package repository

import (
	sqlc "TimeManagementSystem/db/sqlc"
)

type TaskRepository interface {
	Create(userId int, task sqlc.Task) (int, error)
	GetByID(userId, taskId int) (sqlc.Task, error)
	GetAll(userId int) ([]sqlc.Task, error)
	Update(userId, taskId int, task sqlc.Task) error
	Delete(userId, taskId int) error
}

type Authorization interface {
	CreateUser(user sqlc.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UserRepository interface {
	Create(user sqlc.User) (int, error)
	GetUserByEmail(email string) (sqlc.User, error)
}
