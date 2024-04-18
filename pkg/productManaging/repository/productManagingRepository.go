package repository

import (
	"github.com/arthit666/shop_api/entities"
	"github.com/arthit666/shop_api/pkg/productManaging/model"
)

type ProductManagingRepository interface {
	Creating(product *entities.Product) (*entities.Product, error)
	Editing(productID uint64, productEditingReq *model.ProductEditingReq) (uint64, error)
	Archiving(productId uint64) error
}
