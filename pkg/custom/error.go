package custom

import (
	"github.com/labstack/echo/v4"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func Error(c echo.Context, statusCode int, message error) error {
	return c.JSON(statusCode, &ErrorMessage{Message: message.Error()})
}
