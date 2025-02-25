package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID `json:"id"`
	OrderDate time.Time `json:"order_date" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	OrderID   string    `json:"order_id"`
	TableID   *string   `json:"table_id" validate:"required"`
}
