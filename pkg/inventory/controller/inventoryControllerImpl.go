package controller

import (
	"net/http"

	custom "github.com/arthit666/shop_api/pkg/custom"
	_inventoryService "github.com/arthit666/shop_api/pkg/inventory/service"
	"github.com/arthit666/shop_api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type inventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
	logger           echo.Logger
}

func NewInventoryControllerImpl(
	inventoryService _inventoryService.InventoryService,
	logger echo.Logger,
) InventoryController {
	return &inventoryControllerImpl{
		inventoryService: inventoryService,
		logger:           logger,
	}
}

func (c *inventoryControllerImpl) Listing(e echo.Context) error {
	playerID, err := validation.CustomerIDGetting(e)
	if err != nil {
		c.logger.Error("Failed to get playerID", err.Error())
		return custom.Error(e, http.StatusUnauthorized, err)
	}

	inventoryListing, err := c.inventoryService.Listing(playerID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, inventoryListing)
}
