package service

import (
	_inventoryModel "github.com/arthit666/shop_api/pkg/inventory/model"
)

type InventoryService interface {
	Listing(customerID string) ([]*_inventoryModel.Inventory, error)
}
