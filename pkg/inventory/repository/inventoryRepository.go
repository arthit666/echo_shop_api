package repository

import (
	entities "github.com/arthit666/shop_api/entities"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Filling(customerID string, itemID uint64, qty int, tx *gorm.DB) ([]*entities.Inventory, error)
	Listing(customerID string) ([]*entities.Inventory, error)
	Removing(customerID string, itemID uint64, limit int, tx *gorm.DB) error
}
