package repository

import (
	"restaurant-management/database"
	"restaurant-management/models"
)

type OrderRepository struct{}

func (u *OrderRepository) CreateOrder(order *models.Order) error {
	return database.DB.Create(order).Error
}

func (u *OrderRepository) GetOrderById(orderID string) (*models.Order, error) {
	var order models.Order
	err := database.DB.Where("order_id = ?", orderID).First(&order).Error
	return &order, err
}

func (u *OrderRepository) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Find(&orders).Error
	return orders, err
}

func (u *OrderRepository) DeleteOrder(orderID string) error {
	return database.DB.Where("order_id = ?", orderID).Delete(&models.Order{}).Error
}
