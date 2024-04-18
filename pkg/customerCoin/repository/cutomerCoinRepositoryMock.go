package repository

import (
	"github.com/arthit666/shop_api/entities"
	"gorm.io/gorm"

	_customerCoinModel "github.com/arthit666/shop_api/pkg/customerCoin/model"

	"github.com/stretchr/testify/mock"
)

type CoinRepositoryMock struct {
	mock.Mock
}

func (m *CoinRepositoryMock) CoinAdding(customerCoinEntity *entities.CustomerCoin, tx *gorm.DB) (*entities.CustomerCoin, error) {
	args := m.Called(customerCoinEntity, tx)
	return args.Get(0).(*entities.CustomerCoin), args.Error(1)
}

func (m *CoinRepositoryMock) Showing(customerID string) (*_customerCoinModel.CustomerCoinShowing, error) {
	args := m.Called(customerID)
	return args.Get(0).(*_customerCoinModel.CustomerCoinShowing), args.Error(1)
}
