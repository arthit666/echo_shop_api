package server

import (
	_customerCoinRepository "github.com/arthit666/shop_api/pkg/customerCoin/repository"
	_inventoryRepository "github.com/arthit666/shop_api/pkg/inventory/repository"
	_productShopController "github.com/arthit666/shop_api/pkg/productShop/controller"
	_productShopRepository "github.com/arthit666/shop_api/pkg/productShop/repository"
	_productShopService "github.com/arthit666/shop_api/pkg/productShop/service"
)

func (s *echoServer) initProductShopRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/product_shop")

	playerCoinRepository := _customerCoinRepository.NewCustomerCoinRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)
	productShopRepository := _productShopRepository.NewProductShopRepositoryImpl(s.db, s.app.Logger)

	productShopService := _productShopService.NewProductShopServiceImpl(
		productShopRepository,
		playerCoinRepository,
		inventoryRepository,
		s.app.Logger,
	)

	productShopController := _productShopController.NewProductControllerImpl(productShopService)

	router.GET("", productShopController.Listing)
	router.POST("/buying", productShopController.Buying, m.CustomerAuthorizing)

}
