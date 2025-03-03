package repository

import (
	"restaurant-management/database"
	"restaurant-management/models"
)

type OrderItemRepository struct{}

func (u *OrderItemRepository) CreateOrderItem(orderItem *models.OrderItem) error {
	return database.DB.Create(orderItem).Error
}

func (u *OrderItemRepository) GetOrderItemById(orderItemId string) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	err := database.DB.Where("order_item_id = ?", orderItemId).First(&orderItem).Error
	return &orderItem, err
}

func (u *OrderItemRepository) GetOrderItems() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem

	err := database.DB.Find(&orderItems).Error
	return orderItems, err

}

func (u *OrderItemRepository) DeleteOrderItem(orderItemId string) error {
	return database.DB.Where("order_item_id = ?", orderItemId).Delete(&models.OrderItem{}).Error
}
