package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	ID          uuid.UUID `json:"id"`
	Quantity    *string   `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	FoodID      *string   `json:"food_id"`
	OrderItemID string    `json:"order_item_id"`
	OrderID     string    `json:"order_id"`
}
