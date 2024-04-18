package tests

import (
	entities "github.com/arthit666/shop_api/entities"
	_customerCoinModel "github.com/arthit666/shop_api/pkg/customerCoin/model"
	_customerCoinRepository "github.com/arthit666/shop_api/pkg/customerCoin/repository"
	_inventoryRepository "github.com/arthit666/shop_api/pkg/inventory/repository"
	_productShop "github.com/arthit666/shop_api/pkg/productShop/exception"
	_productShopModel "github.com/arthit666/shop_api/pkg/productShop/model"
	_productShopRepository "github.com/arthit666/shop_api/pkg/productShop/repository"
	_productShopService "github.com/arthit666/shop_api/pkg/productShop/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"testing"
)

func TestItemBuyingSuccess(t *testing.T) {
	productShopRepositoryMock := new(_productShopRepository.ProductShopRepositoryMock)
	customerCoinRepositoryMock := new(_customerCoinRepository.CoinRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)
	echoLogger := echo.New().Logger

	itemShopService := _productShopService.NewProductShopServiceImpl(
		productShopRepositoryMock,
		customerCoinRepositoryMock,
		inventoryRepositoryMock,
		echoLogger,
	)

	tx := &gorm.DB{}
	productShopRepositoryMock.On("BeginTransaction").Return(tx)
	productShopRepositoryMock.On("CommitTransaction", tx).Return(nil)
	productShopRepositoryMock.On("RollbackTransaction", tx).Return(nil)

	productShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Product{
		ID:          1,
		Name:        "Product Tester",
		Price:       1000,
		Description: "des test",
		Picture:     "https://www.google.com/tester.jpg",
	}, nil)

	customerCoinRepositoryMock.On("Showing", "P001").Return(&_customerCoinModel.CustomerCoinShowing{
		CustomerID: "P001",
		Coin:       5000,
	}, nil)

	productShopRepositoryMock.On("PurchaseHistoryRecording", &entities.PurchaseHistory{
		CustomerID:         "P001",
		ProductID:          1,
		ProductName:        "Product Tester",
		ProductDescription: "des test",
		ProductPicture:     "https://www.google.com/tester.jpg",
		ProductPrice:       1000,
		Quantity:           3,
		IsBuying:           true,
	}, tx).Return(&entities.PurchaseHistory{
		CustomerID:         "P001",
		ProductID:          1,
		ProductName:        "Product Tester",
		ProductDescription: "des test",
		ProductPicture:     "https://www.google.com/tester.jpg",
		ProductPrice:       1000,
		Quantity:           3,
		IsBuying:           true,
	}, nil)

	inventoryRepositoryMock.On("Filling", "P001", uint64(1), int(3), tx).Return([]*entities.Inventory{
		{
			CustomerID: "P001",
			ProductID:  1,
		},
		{
			CustomerID: "P001",
			ProductID:  1,
		},
		{
			CustomerID: "P001",
			ProductID:  1,
		},
	}, nil)

	customerCoinRepositoryMock.On("CoinAdding", &entities.CustomerCoin{
		CustomerID: "P001",
		Amount:     -3000,
	}, tx).Return(&entities.CustomerCoin{
		CustomerID: "P001",
		Amount:     -3000,
	}, nil)

	type args struct {
		label    string
		in       *_productShopModel.BuyingReq
		expected *_customerCoinModel.CustomerCoin
	}

	cases := []args{
		{
			label: "Success buying item",
			in: &_productShopModel.BuyingReq{
				CustomerID: "P001",
				ProductID:  1,
				Quantity:   3,
			},
			expected: &_customerCoinModel.CustomerCoin{
				CustomerID: "P001",
				Amount:     -3000,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Buying(c.in)
			assert.NoError(t, err)
			assert.EqualValues(t, c.expected, result)
		})
	}
}

func TestItemBuyingFail(t *testing.T) {
	productShopRepositoryMock := new(_productShopRepository.ProductShopRepositoryMock)
	inventoryRepositoryMock := new(_inventoryRepository.InventoryRepositoryMock)
	customerCoinRepositoryMock := new(_customerCoinRepository.CoinRepositoryMock)
	echoLogger := echo.New().Logger

	itemShopService := _productShopService.NewProductShopServiceImpl(
		productShopRepositoryMock,
		customerCoinRepositoryMock,
		inventoryRepositoryMock,
		echoLogger,
	)

	tx := &gorm.DB{}
	productShopRepositoryMock.On("BeginTransaction").Return(tx)
	productShopRepositoryMock.On("CommitTransaction", tx).Return(nil)
	productShopRepositoryMock.On("RollbackTransaction", tx).Return(nil)

	productShopRepositoryMock.On("FindByID", uint64(1)).Return(&entities.Product{
		ID:          1,
		Name:        "Product Tester",
		Price:       1000,
		Description: "des test",
		Picture:     "https://www.google.com/tester.jpg",
	}, nil)

	customerCoinRepositoryMock.On("Showing", "P001").Return(&_customerCoinModel.CustomerCoinShowing{
		CustomerID: "P001",
		Coin:       2000,
	}, nil)

	type args struct {
		label    string
		in       *_productShopModel.BuyingReq
		expected error
	}

	cases := []args{
		{
			label: "Test failed to find item 1",
			in: &_productShopModel.BuyingReq{
				CustomerID: "P001",
				ProductID:  1,
				Quantity:   3,
			},
			expected: &_productShop.CoinNotEnough{},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Buying(c.in)
			assert.Nil(t, result)
			assert.Error(t, err)
			assert.EqualValues(t, c.expected, err)
		})
	}
}
