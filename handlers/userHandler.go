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

var userRepo = repository.UserRepository{}

func CreateUser(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		ID:           uuid.New(),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Password:     req.Password,
		Email:        req.Email,
		Avatar:       req.Avatar,
		Phone:        req.Phone,
		Token:        req.Token,
		RefreshToken: req.RefreshToken,
		CreatedAt:    time.Now(),
		UserId:       req.UserId,
	}

	if err := userRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		log.Printf("Failed to create user: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfull", "user": user})
}

func GetUserById(c *gin.Context) {
	id := c.Param("user_id")
	users, err := userRepo.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
	}
	c.JSON(http.StatusOK, users)
}

func GetUsers(c *gin.Context) {
	users, err := userRepo.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("user_id")

	if err := userRepo.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
