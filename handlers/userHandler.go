package handlers

import (
	"fmt"
	"log"
	"net/http"
	"restaurant-management/helpers"
	"restaurant-management/models"
	"restaurant-management/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.User

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		count, err := userRepo.CheckUser(*req.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
			return
		}
		password := HashPassword(*req.Password)
		req.Password = &password

		count, err = userRepo.CheckPhone(*req.Phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while checking for the phone number"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email and phone number already exist"})
			return
		}

		token, refreshToken, _ := helpers.GenerateAllTokens(*req.Email, *req.FirstName, *req.LastName, req.UserId)
		req.Token = &token
		req.RefreshToken = &refreshToken

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

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var foundUser models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := userRepo.GetUserByEmail(*user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, foundUser.UserId)

		_ = token
		_ = refreshToken
		//helpers.UpdateAllTokens(token, refreshToken, foundUser.UserId)

		c.JSON(http.StatusOK, foundUser)
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providePassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or password is incorrect")
		check = false
	}
	return check, msg
}
