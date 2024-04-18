package controller

import "github.com/labstack/echo/v4"

type ProductManagingController interface {
	Creating(e echo.Context) error
	Editing(e echo.Context) error
	Archiving(e echo.Context) error
}
