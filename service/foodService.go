package service

import (
	"errors"
	"restaurant-management/database"
	"restaurant-management/models"

	"gorm.io/gorm"
)

type FoodService struct {
	db gorm.DB
}

func NewFoodService() *FoodService {
	foodService := new(FoodService)
	foodService.db = *database.GetDB()
	return foodService
}

func (c *FoodService) GetFood(foodId string) (*models.Food, error) {
	class := &models.Food{}
	err := c.db.Model(&models.Food{}).First(class, "id=?", foodId).Error
	if err != nil {
		return nil, errors.New("can't get class ID")
	}
	return class, nil
}
