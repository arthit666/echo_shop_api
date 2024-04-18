package controller

import (
	"net/http"
	"strconv"

	"github.com/arthit666/shop_api/pkg/custom"
	"github.com/arthit666/shop_api/pkg/productManaging/model"
	"github.com/arthit666/shop_api/pkg/productManaging/service"
	"github.com/arthit666/shop_api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type productManagingControllerImpl struct {
	ManagingService service.ProductManagingService
}

func NewProductManagingController(service service.ProductManagingService) ProductManagingController {
	return &productManagingControllerImpl{ManagingService: service}
}

func (c *productManagingControllerImpl) Creating(e echo.Context) error {
	adminID, err := validation.AdminIDGetting(e)
	if err != nil {
		return custom.Error(e, http.StatusUnauthorized, err)
	}

	itemCreatingReq := &model.ProductCreatingReq{}

	validatingContext := custom.NewCustomEchoRequest(e)

	if err := validatingContext.Bind(itemCreatingReq); err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}
	itemCreatingReq.AdminID = adminID

	item, err := c.ManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(e, http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusCreated, item)
}

func (c *productManagingControllerImpl) Editing(e echo.Context) error {
	adminID, err := validation.AdminIDGetting(e)
	if err != nil {
		return custom.Error(e, http.StatusUnauthorized, err)
	}

	itemID, err := c.getItemID(e)
	if err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}

	editItemReq := &model.ProductEditingReq{}

	validatingContext := custom.NewCustomEchoRequest(e)
	if err := validatingContext.Bind(editItemReq); err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}
	editItemReq.AdminID = adminID

	item, err := c.ManagingService.Editing(itemID, editItemReq)
	if err != nil {
		return custom.Error(e, http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, item)
}

func (c *productManagingControllerImpl) Archiving(e echo.Context) error {
	_, err := validation.AdminIDGetting(e)
	if err != nil {
		return custom.Error(e, http.StatusUnauthorized, err)
	}

	itemID, err := c.getItemID(e)
	if err != nil {
		return custom.Error(e, http.StatusBadRequest, err)
	}

	err = c.ManagingService.Archiving(itemID)
	if err != nil {
		return custom.Error(e, http.StatusInternalServerError, err)
	}

	return e.NoContent(http.StatusNoContent)
}

func (c *productManagingControllerImpl) getItemID(pctx echo.Context) (uint64, error) {
	productID := pctx.Param("productId")
	productIDUint64, err := strconv.ParseUint(productID, 10, 64)
	if err != nil {
		return 0, err
	}

	return productIDUint64, nil
}
