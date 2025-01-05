package repositories

import (
	"webserver/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}
