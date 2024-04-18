package controller

import (
	"net/http"

	"github.com/arthit666/shop_api/pkg/custom"
	"github.com/arthit666/shop_api/pkg/productShop/model"
	"github.com/arthit666/shop_api/pkg/productShop/service"
	"github.com/arthit666/shop_api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type productShopControllerImpl struct {
	productShopService service.ProductShopService
}

func NewProductControllerImpl(service service.ProductShopService) ProductShopController {
	return &productShopControllerImpl{productShopService: service}
}

func (c *productShopControllerImpl) Listing(e echo.Context) error {
	productFilter := model.ProductFilter{}

	validatingContext := custom.NewCustomEchoRequest(e)

	if err := validatingContext.Bind(&productFilter); err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}

	productModelList, err := c.productShopService.Listing(&productFilter)
	if err != nil {
		return custom.Error(e, http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, productModelList)
}

func (c *productShopControllerImpl) Buying(e echo.Context) error {
	customerID, err := validation.CustomerIDGetting(e)
	if err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}

	buyingReq := &model.BuyingReq{}

	validatingContext := custom.NewCustomEchoRequest(e)

	if err := validatingContext.Bind(buyingReq); err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}
	buyingReq.CustomerID = customerID

	result, err := c.productShopService.Buying(buyingReq)
	if err != nil {
		return custom.Error(e, http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, result)
}
