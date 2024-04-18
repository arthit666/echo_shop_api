package service

import (
	_customerCoinModel "github.com/arthit666/shop_api/pkg/customerCoin/model"
)

type PlayerCoinService interface {
	CoinAdding(coinAddingReq *_customerCoinModel.CoinAddingReq) (*_customerCoinModel.CustomerCoin, error)
	CustomerCoinShowing(customerID string) *_customerCoinModel.CustomerCoinShowing
}
