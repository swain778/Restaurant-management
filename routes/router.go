package routes

import (
	"restaurant-management/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/foods", handlers.CreateFood)
	r.GET("/foods/food_id", handlers.GetFoodById)
	r.GET("/foods", handlers.GetFoods)
	r.DELETE("/foods/food_id", handlers.DeleteFood)

	return r
}
