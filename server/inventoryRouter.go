package server

import (
	_inventoryController "github.com/arthit666/shop_api/pkg/inventory/controller"
	_inventoryRepository "github.com/arthit666/shop_api/pkg/inventory/repository"
	_inventoryService "github.com/arthit666/shop_api/pkg/inventory/service"
	_productShopRepository "github.com/arthit666/shop_api/pkg/productShop/repository"
)

func (s *echoServer) initInventoryRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/inventory")

	itemRepository := _productShopRepository.NewProductShopRepositoryImpl(s.db, s.app.Logger)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db, s.app.Logger)

	inventoryService := _inventoryService.NewInventoryServiceImpl(
		inventoryRepository,
		itemRepository,
	)

	inventoryController := _inventoryController.NewInventoryControllerImpl(inventoryService, s.app.Logger)

	router.GET("", inventoryController.Listing, m.CustomerAuthorizing)
}
