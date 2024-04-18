package service

import (
	entities "github.com/arthit666/shop_api/entities"
	_customerCoinModel "github.com/arthit666/shop_api/pkg/customerCoin/model"
	_customerCoinRepository "github.com/arthit666/shop_api/pkg/customerCoin/repository"
)

type playerCoinServiceImpl struct {
	customerCoinRepository _customerCoinRepository.CustomerCoinRepository
}

func NewCustomerCoinServiceImpl(
	playerCoinRepository _customerCoinRepository.CustomerCoinRepository,
) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepository}
}

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *_customerCoinModel.CoinAddingReq) (*_customerCoinModel.CustomerCoin, error) {
	playerCoinEntity := &entities.CustomerCoin{
		CustomerID: coinAddingReq.CustomerID,
		Amount:     coinAddingReq.Amount,
	}

	playerCoin, err := s.customerCoinRepository.CoinAdding(playerCoinEntity, nil)
	if err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinModel(), nil
}

func (s *playerCoinServiceImpl) CustomerCoinShowing(customerID string) *_customerCoinModel.CustomerCoinShowing {
	coin, err := s.customerCoinRepository.Showing(customerID)
	if err != nil {
		return &_customerCoinModel.CustomerCoinShowing{
			CustomerID: customerID,
			Coin:       0,
		}
	}

	return coin
}
