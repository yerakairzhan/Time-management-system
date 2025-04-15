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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// логика генерации токена
	return "example-token", nil
}

func (s *AuthService) ParseToken(token string) (int, error) {
	// логика парсинга токена
	return 1, nil
}
