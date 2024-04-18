package controller

import "github.com/labstack/echo/v4"

type InventoryController interface {
	Listing(e echo.Context) error
}
