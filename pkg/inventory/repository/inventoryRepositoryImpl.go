package repository

import (
	"github.com/arthit666/shop_api/databases"
	entities "github.com/arthit666/shop_api/entities"
	_inventory "github.com/arthit666/shop_api/pkg/inventory/exception"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type inventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) InventoryRepository {
	return &inventoryRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *inventoryRepositoryImpl) Filling(customerID string, productID uint64, qty int, tx *gorm.DB) ([]*entities.Inventory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntities := []*entities.Inventory{}

	for range qty {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			CustomerID: customerID,
			ProductID:  productID,
		})
	}

	if err := conn.Create(inventoryEntities).Error; err != nil {
		r.logger.Error("Filling inventory failed:", err.Error())
		return nil, &_inventory.InventoryFilling{
			CustomerID: customerID,
			ProductID:  productID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Listing(customerID string) ([]*entities.Inventory, error) {
	inventoryEntities := []*entities.Inventory{}

	if err := r.db.Connect().Where(
		"customer_id = ? and is_deleted = ?", customerID, false,
	).Find(&inventoryEntities).Error; err != nil {
		r.logger.Error("Listing customer's product failed:", err.Error())
		return nil, &_inventory.CustomerItemsFinding{
			CustomerID: customerID,
		}
	}

	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Removing(cutomerID string, productID uint64, limit int, tx *gorm.DB) error {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	inventoryEntities, err := r.findCustomerItemInInventoryByID(cutomerID, productID, limit)
	if err != nil {
		return err
	}

	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true

		if err := conn.Model(
			&entities.Inventory{},
		).Where(
			"id = ?", inventory.ID,
		).Updates(
			inventory,
		).Error; err != nil {
			r.logger.Error("Removing product failed:", err.Error())
			return &_inventory.CustomerItemRemoving{ProductID: productID}
		}
	}

	return nil
}

func (r *inventoryRepositoryImpl) findCustomerItemInInventoryByID(
	customerID string,
	productID uint64,
	limit int,
) ([]*entities.Inventory, error) {
	inventoryEntities := []*entities.Inventory{}

	if err := r.db.Connect().Where(
		"customer_id = ? and product_id = ? and is_deleted = ?", customerID, productID, false,
	).Limit(
		limit,
	).Find(&inventoryEntities).Error; err != nil {
		r.logger.Error("Finding player's item in inventory failed:", err.Error())
		return nil, &_inventory.CustomerItemsFinding{CustomerID: customerID}
	}

	return inventoryEntities, nil
}
