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

var invoiceRepo = repository.InvoiceRepository{}

func CreateInvoice(c *gin.Context) {
	var req models.Invoice

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice := models.Invoice{
		ID:             uuid.New(),
		InvoiceID:      req.InvoiceID,
		OrderID:        req.OrderID,
		PaymentMethod:  req.PaymentMethod,
		PaymentStatus:  req.PaymentStatus,
		PaymentDueDate: req.PaymentDueDate,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := invoiceRepo.CreateInvoice(&invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
		log.Printf("Failed to create invoice: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Invoice created successfull", "invoice": invoice})
}

func GetInvoiceById(c *gin.Context) {
	id := c.Param("invoice_id")

	invoice, err := invoiceRepo.GetInvoiceById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func GetInvoices(c *gin.Context) {
	invoices, err := invoiceRepo.GetInvoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invoices"})
		return
	}
	c.JSON(http.StatusOK, invoices)
}

func DeleteInvoice(c *gin.Context) {
	id := c.Param("invoice_id")

	if err := invoiceRepo.DeleteInvoice(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete invoice"})
		log.Printf("Failed to delete invoice: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully"})
}
