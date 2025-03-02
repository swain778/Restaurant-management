package repository

import (
	"restaurant-management/database"
	"restaurant-management/models"
)

type FoodRepository struct{}

func (u *FoodRepository) CreateFood(user *models.Food) error {
	return database.DB.Create(user).Error
}

func (u *FoodRepository) GetFoodById(foodID string) (*models.Food, error) {
	var food models.Food
	err := database.DB.Where("food_id = ?", foodID).First(&food).Error
	return &food, err
}

func (u *FoodRepository) GetFoods() ([]models.Food, error) {
	var foods []models.Food

	err := database.DB.Find(&foods).Error
	return foods, err
}

func (u *FoodRepository) DeleteFood(foodID string) error {
	return database.DB.Where("food_id = ?", foodID).Delete(&models.Food{}).Error
}
