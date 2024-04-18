package service

import (
	"github.com/arthit666/shop_api/entities"
	_customerCoin "github.com/arthit666/shop_api/pkg/customerCoin/model"
	_customerCoinRepo "github.com/arthit666/shop_api/pkg/customerCoin/repository"
	_itemShopException "github.com/arthit666/shop_api/pkg/productShop/exception"

	_inventoryRepository "github.com/arthit666/shop_api/pkg/inventory/repository"
	_productShop "github.com/arthit666/shop_api/pkg/productShop/model"
	_productShopRepo "github.com/arthit666/shop_api/pkg/productShop/repository"
	"github.com/labstack/echo/v4"
)

type productShopServiceImpl struct {
	productShopRepository  _productShopRepo.ProductShopRepository
	customerCoinRepository _customerCoinRepo.CustomerCoinRepository
	inventoryRepository    _inventoryRepository.InventoryRepository
	logger                 echo.Logger
}

func NewProductShopServiceImpl(
	productShopRepository _productShopRepo.ProductShopRepository,
	customerCoinRepository _customerCoinRepo.CustomerCoinRepository,
	inventoryRepository _inventoryRepository.InventoryRepository,
	logger echo.Logger,
) ProductShopService {
	return &productShopServiceImpl{productShopRepository, customerCoinRepository, inventoryRepository, logger}
}

func (s *productShopServiceImpl) Listing(productFilter *_productShop.ProductFilter) (*_productShop.ProductResult, error) {
	productList, err := s.productShopRepository.Listing(productFilter)
	if err != nil {
		return nil, err
	}

	productCounting, err := s.productShopRepository.Counting(productFilter)
	if err != nil {
		return nil, err
	}

	totalPage := s.totalPageCalculation(productCounting, productFilter.Size)
	productResuilt := s.toProductResultsResponse(productList, productFilter.Page, totalPage)

	return productResuilt, nil
}

func (s *productShopServiceImpl) totalPageCalculation(totalItems, size int64) int64 {
	totalPage := totalItems / size

	if totalItems%size != 0 {
		totalPage++
	}

	return totalPage
}

func (s *productShopServiceImpl) toProductResultsResponse(productEntityList []*entities.Product, page, totalPage int64) *_productShop.ProductResult {
	products := []*_productShop.Product{}

	for _, itemEntity := range productEntityList {
		products = append(products, itemEntity.ToProductModel())
	}

	return &_productShop.ProductResult{
		Products: products,
		Paginate: _productShop.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}

func (s *productShopServiceImpl) Buying(buyingReq *_productShop.BuyingReq) (*_customerCoin.CustomerCoin, error) {
	productEntity, err := s.productShopRepository.FindByID(buyingReq.ProductID)
	if err != nil {
		return nil, err
	}

	totalPrice := s.calculateTotalPrice(productEntity.ToProductModel(), buyingReq.Quantity)

	if err := s.customerCoinChecking(buyingReq.CustomerID, totalPrice); err != nil {
		return nil, err
	}

	tx := s.productShopRepository.BeginTransaction()

	purchaseRecording, err := s.productShopRepository.PurchaseHistoryRecording(&entities.PurchaseHistory{
		CustomerID:         buyingReq.CustomerID,
		ProductID:          productEntity.ID,
		ProductName:        productEntity.Name,
		ProductDescription: productEntity.Description,
		ProductPrice:       productEntity.Price,
		ProductPicture:     productEntity.Picture,
		Quantity:           buyingReq.Quantity,
		IsBuying:           true,
	}, tx)
	if err != nil {
		s.productShopRepository.RollbackTransaction(tx)
		return nil, err
	}
	s.logger.Infof("Purchase history recorded: %d", purchaseRecording.ID)

	coinRecording, err := s.customerCoinRepository.CoinAdding(&entities.CustomerCoin{
		CustomerID: buyingReq.CustomerID,
		Amount:     -totalPrice,
	}, tx)
	if err != nil {
		s.productShopRepository.RollbackTransaction(tx)
		return nil, err
	}
	s.logger.Infof("Player coins reduced for: %d", totalPrice)

	inventoryRecording, err := s.inventoryRepository.Filling(
		buyingReq.CustomerID,
		buyingReq.ProductID,
		int(buyingReq.Quantity),
		tx,
	)
	if err != nil {
		s.productShopRepository.RollbackTransaction(tx)
		return nil, err
	}
	s.logger.Infof("Items recorded into player inventory: %d", len(inventoryRecording))

	if err := s.productShopRepository.CommitTransaction(tx); err != nil {
		s.productShopRepository.RollbackTransaction(tx)
		return nil, err
	}

	return coinRecording.ToPlayerCoinModel(), nil
}

func (s *productShopServiceImpl) calculateTotalPrice(product *_productShop.Product, qty uint) int64 {
	return int64(product.Price) * int64(qty)
}

func (s *productShopServiceImpl) customerCoinChecking(playerID string, totalPrice int64) error {
	playerCoin, err := s.customerCoinRepository.Showing(playerID)
	if err != nil {
		return err
	}

	if playerCoin.Coin < totalPrice {
		s.logger.Errorf("Customer %s has not enough coin", playerID)
		return &_itemShopException.CoinNotEnough{}
	}

	return nil
}
