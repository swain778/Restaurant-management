package handlers

import (
	"log"
	"net/http"
	"restaurant-management/models"
	"restaurant-management/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var foodRepo = repository.FoodRepository{}

func CreateFood(c *gin.Context) {
	var req models.Food

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	food := models.Food{
		ID:        uuid.New(),
		Name:      req.Name,
		Price:     req.Price,
		FoodImage: req.FoodImage,
		FoodID:    req.FoodID,
		MenuID:    req.MenuID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := foodRepo.CreateFood(&food); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		log.Printf("Failed to create user: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfull", "user": food})
}

func GetFoodById(c *gin.Context) {
	id := c.Param("food_id")

	food, err := foodRepo.GetFoodById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, food)
}

func GetFoods(c *gin.Context) {
	foods, err := foodRepo.GetFoods()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, foods)
}

func DeleteFood(c *gin.Context) {
	id := c.Param("food_id")

	if err := foodRepo.DeleteFood(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
}
