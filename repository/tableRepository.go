package repository

import (
	"restaurant-management/database"
	"restaurant-management/models"
)

type TableRepository struct{}

func (u *TableRepository) CreateTable(table *models.Table) error {
	return database.DB.Create(table).Error
}

func (u *TableRepository) GetTableById(tableID string) (*models.Table, error) {
	var table models.Table
	err := database.DB.Where("table_id = ?", tableID).First(&table).Error
	return &table, err
}

func (u *TableRepository) GetTables() ([]models.Table, error) {
	var tables []models.Table

	err := database.DB.Find(&tables).Error
	return tables, err
}

func (u *TableRepository) DeleteTable(tableID string) error {
	return database.DB.Where("table_id = ?", tableID).Delete(&models.Table{}).Error
}
