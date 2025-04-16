package service

import (
	db "TimeManagementSystem/db/sqlc"
	"TimeManagementSystem/repository"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user db.User) (int, error) {
	return s.repo.Create(user)
}

func (s *AuthService) GetUserByEmail(email string) (db.User, error) {
	return s.repo.GetUserByEmail(email)
}
