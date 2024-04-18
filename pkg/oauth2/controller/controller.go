package controller

import "github.com/labstack/echo/v4"

type OAuth2Controller interface {
	CustomerLogin(e echo.Context) error
	AdminLogin(e echo.Context) error
	CustomerLoginCallback(e echo.Context) error
	AdminLoginCallback(e echo.Context) error
	Logout(e echo.Context) error

	CustomerAuthorizing(e echo.Context, next echo.HandlerFunc) error
	AdminAuthorizing(e echo.Context, next echo.HandlerFunc) error
}
