package repository

import (
	"restaurant-management/database"
	"restaurant-management/models"
)

type FoodRepository struct{}

func (u *FoodRepository) CreateFood(user *models.Food) error {
	return database.DB.Create(user).Error
}

func (u *FoodRepository) GetFoodById(foodID uint) (*models.Food, error) {
	var food models.Food
	err := database.DB.First(&food, foodID).Error
	return &food, err
}

func (u *FoodRepository) GetFoods() ([]models.Food, error) {
	var foods []models.Food

	err := database.DB.Find(&foods).Error
	return foods, err
}

func (u *FoodRepository) DeleteFood(foodID uint) error {
	return database.DB.Delete(&models.Food{}, foodID).Error
}
