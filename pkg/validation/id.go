package validation

import (
	"github.com/labstack/echo/v4"

	_admin "github.com/arthit666/shop_api/pkg/admin/exception"
	_customer "github.com/arthit666/shop_api/pkg/customer/exception"
)

func AdminIDGetting(e echo.Context) (string, error) {
	if adminID, ok := e.Get("adminID").(string); !ok || adminID == "" {
		return "", &_admin.AdminNotFound{AdminID: "Unknown"}
	} else {
		return adminID, nil
	}
}

func CustomerIDGetting(e echo.Context) (string, error) {
	if playerID, ok := e.Get("playerID").(string); !ok || playerID == "" {
		return "", &_customer.CustomerNotFound{CustomID: "Unknown"}
	} else {
		return playerID, nil
	}
}
