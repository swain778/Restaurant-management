package models

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name" validate:"required"`
	Category  string     `json:"category" validate:"required"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	MenuID    string     `json:"food_id`
}
