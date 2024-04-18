package repository

import (
	"github.com/arthit666/shop_api/databases"
	"github.com/arthit666/shop_api/entities"
	"github.com/arthit666/shop_api/pkg/productShop/exception"
	"github.com/arthit666/shop_api/pkg/productShop/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type productShopRepositoryImpl struct {
	logger echo.Logger
	db     databases.Database
}

func NewProductShopRepositoryImpl(db databases.Database, logger echo.Logger) ProductShopRepository {
	return &productShopRepositoryImpl{db: db, logger: logger}
}

func (r *productShopRepositoryImpl) Listing(productFilter *model.ProductFilter) ([]*entities.Product, error) {

	query := r.db.Connect().Model(&entities.Product{}).Where("is_deleted = ?", false)

	if productFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+productFilter.Name+"%")
	}
	if productFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+productFilter.Description+"%")
	}

	offset := int((productFilter.Page - 1) * productFilter.Size)
	size := int(productFilter.Size)

	productList := []*entities.Product{}

	if err := query.Offset(offset).Limit(size).Find(&productList).Order("id asc").Error; err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, &exception.ProductListing{}
	}

	return productList, nil
}

func (r *productShopRepositoryImpl) Counting(productFilter *model.ProductFilter) (int64, error) {
	query := r.db.Connect().Model(&entities.Product{}).Where("is_deleted = ?", false)

	if productFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+productFilter.Name+"%")
	}
	if productFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+productFilter.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Error("Counting items failed:", err.Error())
		return -1, &exception.ProductCounting{}
	}

	return count, nil
}

func (r *productShopRepositoryImpl) FindByID(productID uint64) (*entities.Product, error) {
	item := &entities.Product{}

	if err := r.db.Connect().First(item, productID).Error; err != nil {
		r.logger.Error("Finding product failed:", err.Error())
		return nil, &exception.ProductNotFound{ProductID: productID}
	}

	return item, nil
}

func (r *productShopRepositoryImpl) FindByIDList(productIDs []uint64) ([]*entities.Product, error) {
	items := []*entities.Product{}

	if err := r.db.Connect().Model(&entities.Product{}).Where("id in ?", productIDs).Find(&items).Error; err != nil {
		r.logger.Error("Finding product by ID failed:", err.Error())
		return nil, &exception.ProductListing{}
	}

	return items, nil
}

func (r *productShopRepositoryImpl) PurchaseHistoryRecording(
	purchasingEntity *entities.PurchaseHistory,
	tx *gorm.DB,
) (*entities.PurchaseHistory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	insertedPurchasing := &entities.PurchaseHistory{}

	if err := conn.Create(purchasingEntity).Scan(insertedPurchasing).Error; err != nil {
		r.logger.Errorf("Purchase history recording failed: %s", err.Error())
		return nil, &exception.HistoryOfPurchaseRecording{}
	}

	return insertedPurchasing, nil
}

func (r *productShopRepositoryImpl) BeginTransaction() *gorm.DB {
	tx := r.db.Connect()
	return tx.Begin()
}

func (r *productShopRepositoryImpl) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *productShopRepositoryImpl) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}
