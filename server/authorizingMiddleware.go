package server

import (
	_oauth2Controller "github.com/arthit666/shop_api/pkg/oauth2/controller"

	"github.com/arthit666/shop_api/config"
	"github.com/labstack/echo/v4"
)

type authorizingMiddleware struct {
	oauth2Controller _oauth2Controller.OAuth2Controller
	oauth2Conf       *config.OAuth2
	logger           echo.Logger
}

func (m *authorizingMiddleware) CustomerAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		return m.oauth2Controller.CustomerAuthorizing(e, next)
	}
}

func (m *authorizingMiddleware) AdminAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		return m.oauth2Controller.AdminAuthorizing(e, next)
	}
}
