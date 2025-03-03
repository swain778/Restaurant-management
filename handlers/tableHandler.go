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

var tableRepo = repository.TableRepository{}

func CreateTable(c *gin.Context) {
	var req models.Table

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	table := models.Table{
		ID:             uuid.New(),
		NumberOfGuests: req.NumberOfGuests,
		TableNumber:    req.TableNumber,
		CreatedAt:      time.Now(),
		TableID:        req.TableID,
	}

	if err := tableRepo.CreateTable(&table); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create table"})
		log.Printf("Failed to create table: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Table created successfully", "table": table})
}

func GetTableById(c *gin.Context) {
	id := c.Param("table_id")
	table, err := tableRepo.GetTableById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
		return
	}
	c.JSON(http.StatusOK, table)
}

func GetTables(c *gin.Context) {
	table, err := tableRepo.GetTables()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tables"})
		return
	}
	c.JSON(http.StatusOK, table)
}

func DeleteTable(c *gin.Context) {
	id := c.Param("table_id")

	if err := tableRepo.DeleteTable(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete table"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Table deleted successfully"})
}
