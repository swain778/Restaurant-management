package database

import (
	"errors"
	"fmt"
	"restaurant-management/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=swain@123 dbname=restaurant_management port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		_ = errors.New("can't connect to database")
	}
	DB = database
	MigrateDB()
}

func MigrateDB() {
	err := DB.Migrator().AutoMigrate(
		&models.Food{},
		&models.Invoice{},
		&models.Menu{},
		&models.Note{},
		&models.OrderItem{},
		&models.Order{},
		&models.Table{},
		&models.User{},
	)
	if err != nil {
		fmt.Println(err.Error())
	}
}
