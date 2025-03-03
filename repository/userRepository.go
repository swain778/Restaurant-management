package repository

import (
	"restaurant-management/database"
	"restaurant-management/models"
)

type UserRepository struct{}

func (u *UserRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func (u *UserRepository) GetUserById(userID string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("user_id = ?", userID).First(&userID).Error
	return &user, err
}

func (u *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User

	err := database.DB.Find(&users).Error
	return users, err
}

func (u *UserRepository) DeleteUser(userID string) error {
	return database.DB.Where("user_id = ?", userID).Delete(&models.User{}).Error
}
