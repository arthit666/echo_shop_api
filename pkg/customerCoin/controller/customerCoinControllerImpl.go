package controller

import (
	"net/http"

	custom "github.com/arthit666/shop_api/pkg/custom"
	_customerCoinModel "github.com/arthit666/shop_api/pkg/customerCoin/model"
	_customerCoinService "github.com/arthit666/shop_api/pkg/customerCoin/service"
	"github.com/arthit666/shop_api/pkg/validation"

	"github.com/labstack/echo/v4"
)

type playerCoinControllerImpl struct {
	customerCoinService _customerCoinService.PlayerCoinService
}

func NewCustomerCoinControllerImpl(customerCoinService _customerCoinService.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{
		customerCoinService: customerCoinService,
	}
}

func (c *playerCoinControllerImpl) CoinAdding(e echo.Context) error {
	adminId, err := validation.AdminIDGetting(e)
	if err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}

	coinAddingReq := &_customerCoinModel.CoinAddingReq{}

	validatingContext := custom.NewCustomEchoRequest(e)

	if err := validatingContext.Bind(coinAddingReq); err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}
	coinAddingReq.CustomerID = adminId

	customerCoin, err := c.customerCoinService.CoinAdding(coinAddingReq)
	if err != nil {
		return custom.Error(e, http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusCreated, customerCoin)
}

func (c *playerCoinControllerImpl) CustomerCoinShowing(e echo.Context) error {
	playerID, err := validation.CustomerIDGetting(e)
	if err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}

	playerCoin := c.customerCoinService.CustomerCoinShowing(playerID)

	return e.JSON(http.StatusOK, playerCoin)

}
