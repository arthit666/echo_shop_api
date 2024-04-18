package repository

import (
	"github.com/arthit666/shop_api/databases"
	"github.com/arthit666/shop_api/entities"
	exception "github.com/arthit666/shop_api/pkg/productManaging/exeption"
	"github.com/arthit666/shop_api/pkg/productManaging/model"
	"github.com/labstack/echo/v4"
)

type productManagingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewProductManagingRepositoryImpl(db databases.Database, logger echo.Logger) ProductManagingRepository {
	return &productManagingRepositoryImpl{db: db, logger: logger}
}

func (r *productManagingRepositoryImpl) Creating(itemEntity *entities.Product) (*entities.Product, error) {
	product := &entities.Product{}

	if err := r.db.Connect().Create(itemEntity).Scan(product).Error; err != nil {
		r.logger.Error("product creating failed:", err.Error())
		return nil, &exception.ProductCreating{}
	}

	return product, nil
}

func (r *productManagingRepositoryImpl) Editing(productID uint64, productEditingReq *model.ProductEditingReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Product{}).Where(
		"id = ?", productID,
	).Updates(
		productEditingReq,
	).Error; err != nil {
		r.logger.Error("Editing product failed:", err.Error())
		return 0, &exception.ProductEditing{}
	}

	return productID, nil
}

func (r *productManagingRepositoryImpl) Archiving(productId uint64) error {
	if err := r.db.Connect().Table("products").Where(
		"id = ?", productId,
	).Update(
		"is_deleted", true,
	).Error; err != nil {
		r.logger.Error("delete product failed:", err.Error())
		return &exception.ProductArchiving{ProductID: productId}
	}

	return nil
}
