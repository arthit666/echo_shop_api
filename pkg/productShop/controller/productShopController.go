package controller

import (
	"github.com/labstack/echo/v4"
)

type ProductShopController interface {
	Listing(c echo.Context) error
	Buying(pctx echo.Context) error
}
