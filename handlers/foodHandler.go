package handlers

import (
	"net/http"
	"restaurant-management/models"
	"restaurant-management/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var foodRepo = repository.FoodRepository{}

func CreateFood(c *gin.Context) {
	var req models.Food

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	food := models.Food{
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
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfull", "user": food})
}

func GetFoodById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("food_id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid food ID"})
		return
	}
	food, err := foodRepo.GetFoodById(uint(id))
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
	id, err := strconv.Atoi(c.Param("food_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := foodRepo.DeleteFood(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
}
