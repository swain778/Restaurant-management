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

var orderRepo = repository.OrderRepository{}

func CreateOrder(c *gin.Context) {
	var req models.Order

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		ID:        uuid.New(),
		OrderDate: time.Now(),
		CreatedAt: time.Now(),
		OrderID:   req.OrderID,
		TableID:   req.TableID,
	}

	if err := orderRepo.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		log.Printf("Failed to create order: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order": order})
}

func GetOrderById(c *gin.Context) {
	id := c.Param("order_id")

	order, err := orderRepo.GetOrderById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
	}
	c.JSON(http.StatusOK, order)
}

func GetOrders(c *gin.Context) {
	orders, err := orderRepo.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("order_id")

	if err := orderRepo.DeleteOrder(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
