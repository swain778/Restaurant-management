package repository

import (
	"restaurant-management/database"
	"restaurant-management/models"
)

type MenuRepository struct{}

func (u *MenuRepository) CreateMenu(menu *models.Menu) error {
	return database.DB.Create(menu).Error
}

func (u *MenuRepository) GetMenuById(menuID string) (*models.Menu, error) {
	var menu models.Menu
	err := database.DB.Where("menu_id = ?", menuID).First(&menu).Error
	return &menu, err
}

func (u *MenuRepository) GetMenus() ([]models.Menu, error) {
	var menus []models.Menu

	err := database.DB.Find(&menus).Error
	return menus, err
}

func (u *MenuRepository) DeleteMenu(menuID string) error {
	return database.DB.Where("menu_id = ?", menuID).Delete(&models.Menu{}).Error
}
