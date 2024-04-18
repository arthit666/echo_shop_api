package service

import (
	entities "github.com/arthit666/shop_api/entities"
	_inventoryModel "github.com/arthit666/shop_api/pkg/inventory/model"
	_inventoryRepository "github.com/arthit666/shop_api/pkg/inventory/repository"
	_productShopRepository "github.com/arthit666/shop_api/pkg/productShop/repository"
)

type inventoryServiceImpl struct {
	inventoryRepository   _inventoryRepository.InventoryRepository
	productShopRepository _productShopRepository.ProductShopRepository
}

func NewInventoryServiceImpl(
	inventoryRepository _inventoryRepository.InventoryRepository,
	ProductShopRepository _productShopRepository.ProductShopRepository,
) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepository,
		ProductShopRepository,
	}
}

func (s *inventoryServiceImpl) Listing(customerID string) ([]*_inventoryModel.Inventory, error) {
	inventories, err := s.inventoryRepository.Listing(customerID)
	if err != nil {
		return nil, err
	}

	uniqueProductCounterList := s.getUniqueProductCounterList(inventories)

	return s.buildInventoryListingResult(
		uniqueProductCounterList,
	), nil
}

func (s *inventoryServiceImpl) buildInventoryListingResult(
	uniqueProductCounterList []_inventoryModel.ProductQuantityCounting,
) []*_inventoryModel.Inventory {
	uniqueProductShopRepositoryIDList := s.getProductIDList(uniqueProductCounterList)

	ProductEntitiesList, err := s.productShopRepository.FindByIDList(uniqueProductShopRepositoryIDList)
	if err != nil {
		return []*_inventoryModel.Inventory{}
	}

	results := []*_inventoryModel.Inventory{}
	ProductMapWithQuantity := s.getProductMapWithQuantity(uniqueProductCounterList)

	for _, ProductEntity := range ProductEntitiesList {
		results = append(results, &_inventoryModel.Inventory{
			Product:  ProductEntity.ToProductModel(),
			Quantity: ProductMapWithQuantity[ProductEntity.ID],
		})
	}

	return results
}

func (s *inventoryServiceImpl) getUniqueProductCounterList(
	inventories []*entities.Inventory,
) []_inventoryModel.ProductQuantityCounting {
	ProductCounterList := []_inventoryModel.ProductQuantityCounting{}

	ProductMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventories {
		ProductMapWithQuantity[inventory.ProductID]++
	}

	for ProductShopRepositoryID, quantity := range ProductMapWithQuantity {
		ProductCounterList = append(ProductCounterList, _inventoryModel.ProductQuantityCounting{
			ProductID: ProductShopRepositoryID,
			Quantity:  quantity,
		})

	}
	return ProductCounterList
}

func (s *inventoryServiceImpl) getProductIDList(
	uniqueProductCounterList []_inventoryModel.ProductQuantityCounting,
) []uint64 {
	uniqueProductIDList := make([]uint64, 0)

	for _, inventory := range uniqueProductCounterList {
		uniqueProductIDList = append(uniqueProductIDList, inventory.ProductID)
	}

	return uniqueProductIDList
}

func (s *inventoryServiceImpl) getProductMapWithQuantity(
	uniqueProductCounterList []_inventoryModel.ProductQuantityCounting,
) map[uint64]uint {
	ProductMapWithQuantity := make(map[uint64]uint)

	for _, ProductQuantityCounter := range uniqueProductCounterList {
		ProductMapWithQuantity[ProductQuantityCounter.ProductID] = ProductQuantityCounter.Quantity
	}

	return ProductMapWithQuantity
}
