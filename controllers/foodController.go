package controllers

import (
	"net/http"
	"restaurant-management/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func GetFoods(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service := service.NewFoodService()
	class, err := service.GetFood(params["id"])
}

func GetFood() gin.HandlerFunc {
	service := service.NewFoodService()

	return func(c *gin.Context) {

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}
