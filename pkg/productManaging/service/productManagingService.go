package service

import (
	_productManagingModel "github.com/arthit666/shop_api/pkg/productManaging/model"
	_productShopModel "github.com/arthit666/shop_api/pkg/productShop/model"
)

type ProductManagingService interface {
	Creating(productCreateReq *_productManagingModel.ProductCreatingReq) (*_productShopModel.Product, error)
	Editing(productID uint64, itemEditingReq *_productManagingModel.ProductEditingReq) (*_productShopModel.Product, error)
	Archiving(itemID uint64) error
}
