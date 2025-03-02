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

var menuRepo = repository.MenuRepository{}

func CreateMenu(c *gin.Context) {
	var req models.Menu

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("Failed to bind JSON: %v", err)
		return
	}

	menu := models.Menu{
		ID:        uuid.New(),
		Name:      req.Name,
		Category:  req.Category,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		CreatedAt: time.Now(),
		MenuID:    req.MenuID,
	}
	if err := menuRepo.CreateMenu(&menu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu"})
		log.Printf("Failed to create menu: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu created successfully", "menu": menu})
}

func GetMenuById(c *gin.Context) {
	id := c.Param("menu_id")

	menu, err := menuRepo.GetMenuById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		log.Printf("Menu not found:%v", err)
		return
	}
	c.JSON(http.StatusOK, menu)
}

func GetMenus(c *gin.Context) {
	foods, err := menuRepo.GetMenus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menus"})
		log.Printf("Failed to fetch menus: %v", err)
		return
	}
	c.JSON(http.StatusOK, foods)
}

func DeleteMenu(c *gin.Context) {
	id := c.Param("menu_id")

	if err := menuRepo.DeleteMenu(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete menu"})
		log.Printf("Failed to delete menu: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}
