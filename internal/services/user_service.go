package services

import (
	"webserver/internal/models"
	"webserver/internal/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.Repo.GetAll()
}
