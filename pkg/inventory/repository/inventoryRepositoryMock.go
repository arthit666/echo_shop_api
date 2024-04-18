package repository

import (
	entities "github.com/arthit666/shop_api/entities"
	"gorm.io/gorm"

	"github.com/stretchr/testify/mock"
)

type InventoryRepositoryMock struct {
	mock.Mock
}

func (m *InventoryRepositoryMock) Filling(customerID string, productID uint64, qty int, tx *gorm.DB) ([]*entities.Inventory, error) {
	args := m.Called(customerID, productID, qty, tx)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) Listing(customerID string) ([]*entities.Inventory, error) {
	args := m.Called(customerID)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}

func (m *InventoryRepositoryMock) Removing(customerID string, productID uint64, limit int, tx *gorm.DB) error {
	args := m.Called(customerID, productID, limit, tx)
	return args.Error(0)
}
