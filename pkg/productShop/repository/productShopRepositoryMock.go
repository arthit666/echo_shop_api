package repository

import (
	entities "github.com/arthit666/shop_api/entities"
	_productShopModel "github.com/arthit666/shop_api/pkg/productShop/model"
	"gorm.io/gorm"

	"github.com/stretchr/testify/mock"
)

type ProductShopRepositoryMock struct {
	mock.Mock
}

func (m *ProductShopRepositoryMock) BeginTransaction() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *ProductShopRepositoryMock) RollbackTransaction(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *ProductShopRepositoryMock) CommitTransaction(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *ProductShopRepositoryMock) FindByID(productID uint64) (*entities.Product, error) {
	args := m.Called(productID)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *ProductShopRepositoryMock) Listing(productFilter *_productShopModel.ProductFilter) ([]*entities.Product, error) {
	args := m.Called(productFilter)
	return args.Get(0).([]*entities.Product), args.Error(1)
}

func (m *ProductShopRepositoryMock) FindByIDList(productIDs []uint64) ([]*entities.Product, error) {
	args := m.Called(productIDs)
	return args.Get(0).([]*entities.Product), args.Error(1)
}

func (m *ProductShopRepositoryMock) Counting(productFilter *_productShopModel.ProductFilter) (int64, error) {
	args := m.Called(productFilter)
	return args.Get(0).(int64), args.Error(1)
}

func (m *ProductShopRepositoryMock) PurchaseHistoryRecording(
	purchasingEntity *entities.PurchaseHistory,
	tx *gorm.DB,
) (*entities.PurchaseHistory, error) {
	args := m.Called(purchasingEntity, tx)
	return args.Get(0).(*entities.PurchaseHistory), args.Error(1)
}
