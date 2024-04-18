package model

import (
	_productShopModel "github.com/arthit666/shop_api/pkg/productShop/model"
)

type (
	Inventory struct {
		Product  *_productShopModel.Product `json:"product"`
		Quantity uint                       `json:"quantity"`
	}

	ProductQuantityCounting struct {
		ProductID uint64
		Quantity  uint
	}
)
