package repository

import (
	"github.com/arthit666/shop_api/entities"
	_customerCoinModel "github.com/arthit666/shop_api/pkg/customerCoin/model"
	"gorm.io/gorm"
)

type CustomerCoinRepository interface {
	CoinAdding(customerCoinEntity *entities.CustomerCoin, tx *gorm.DB) (*entities.CustomerCoin, error)
	Showing(customerID string) (*_customerCoinModel.CustomerCoinShowing, error)
}
