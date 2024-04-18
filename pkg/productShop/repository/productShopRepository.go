package repository

import (
	"github.com/arthit666/shop_api/entities"
	"github.com/arthit666/shop_api/pkg/productShop/model"
	"gorm.io/gorm"
)

type ProductShopRepository interface {
	BeginTransaction() *gorm.DB
	RollbackTransaction(tx *gorm.DB) error
	CommitTransaction(tx *gorm.DB) error
	Listing(productFilter *model.ProductFilter) ([]*entities.Product, error)
	Counting(productFilter *model.ProductFilter) (int64, error)
	FindByID(productID uint64) (*entities.Product, error)
	FindByIDList(productIDs []uint64) ([]*entities.Product, error)
	PurchaseHistoryRecording(
		purchasingEntity *entities.PurchaseHistory,
		tx *gorm.DB,
	) (*entities.PurchaseHistory, error)
}
