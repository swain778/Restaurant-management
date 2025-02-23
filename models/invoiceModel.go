package models

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID             uuid.UUID `json:"id"`
	InvoiceID      string    `json:"invoice_id"`
	OrderID        string    `json:"order_id"`
	PaymentMethod  *string   `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	PaymentStatus  *string   `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_date"`
}
