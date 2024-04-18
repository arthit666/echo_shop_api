package service

import (
	_customerCoin "github.com/arthit666/shop_api/pkg/customerCoin/model"
	"github.com/arthit666/shop_api/pkg/productShop/model"
)

type ProductShopService interface {
	Listing(productFilter *model.ProductFilter) (*model.ProductResult, error)
	Buying(buyingReq *model.BuyingReq) (*_customerCoin.CustomerCoin, error)
}
