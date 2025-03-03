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

var orderItemRepo = repository.OrderItemRepository{}

func CreateOrderItem(c *gin.Context) {
	var req models.OrderItem

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItem := models.OrderItem{
		ID:          uuid.New(),
		Quantity:    req.Quantity,
		CreatedAt:   time.Now(),
		FoodID:      req.FoodID,
		OrderItemID: req.OrderItemID,
		OrderID:     req.OrderID,
	}
	if err := orderItemRepo.CreateOrderItem(&orderItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order item"})
		log.Printf("Failed to create order item: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order item created successfully", "orderItem": orderItem})
}

func GetOrderItemById(c *gin.Context) {
	id := c.Param("order_item_id")

	orderItem, err := orderItemRepo.GetOrderItemById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func OrderItems(c *gin.Context) {
	orderItem, err := orderItemRepo.GetOrderItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order items"})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func DeleteOrderItem(c *gin.Context) {
	id := c.Param("order_item_id")
	err := orderItemRepo.DeleteOrderItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
}
