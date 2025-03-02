package repository

import (
	"restaurant-management/database"
	"restaurant-management/models"
)

type InvoiceRepository struct{}

func (u *InvoiceRepository) CreateInvoice(invoice *models.Invoice) error {
	return database.DB.Create(invoice).Error
}

func (u *InvoiceRepository) GetInvoiceById(invoiceID string) (*models.Invoice, error) {
	var invoice models.Invoice
	err := database.DB.Where("invoice_id = ?", invoiceID).First(&invoice).Error
	return &invoice, err
}

func (u *InvoiceRepository) GetInvoices() ([]models.Invoice, error) {
	var invoices []models.Invoice

	err := database.DB.Find(&invoices).Error
	return invoices, err
}

func (u *InvoiceRepository) DeleteInvoice(invoiceID string) error {
	return database.DB.Where("invoice_id = ?", invoiceID).Delete(&models.Invoice{}).Error
}
