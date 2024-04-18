package repository

import (
	"github.com/arthit666/shop_api/databases"
	"github.com/arthit666/shop_api/entities"
	_customerCoin "github.com/arthit666/shop_api/pkg/customerCoin/exception"
	_customerCoinModel "github.com/arthit666/shop_api/pkg/customerCoin/model"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewCustomerCoinRepositoryImpl(db databases.Database, logger echo.Logger) CustomerCoinRepository {
	return &playerCoinRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerCoinRepositoryImpl) CoinAdding(customerCoinEntity *entities.CustomerCoin, tx *gorm.DB) (*entities.CustomerCoin, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	customerCoin := &entities.CustomerCoin{}

	if err := conn.Create(customerCoinEntity).Scan(customerCoin).Error; err != nil {
		r.logger.Error("Player's balance recording failed:", err.Error())
		return nil, &_customerCoin.CoinAdding{}
	}

	return customerCoin, nil
}

func (r *playerCoinRepositoryImpl) Showing(customerID string) (*_customerCoinModel.CustomerCoinShowing, error) {
	customerCoin := &_customerCoinModel.CustomerCoinShowing{}

	if err := r.db.Connect().Model(
		&entities.CustomerCoin{},
	).Where(
		"customer_id = ?", customerID,
	).Select(
		"customer_id, sum(amount) as coin",
	).Group(
		"customer_id",
	).Scan(&customerCoin).Error; err != nil {
		r.logger.Error("Calculating player coin failed:", err.Error())
		return nil, &_customerCoin.CustomerCoinShowing{}
	}

	return customerCoin, nil
}
