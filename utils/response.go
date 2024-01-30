package utils

import (
	"github.com/labstack/echo/v4"
)

func SuccessResponse(c echo.Context, statusCode int, data interface{}) error {
	return c.JSON(statusCode, map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}

func InfoResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, map[string]string{
		"message": message,
		"status":  "success",
	})
}

func ErrorResponse(c echo.Context, statusCode int, errors []string) error {
	return c.JSON(statusCode, map[string]interface{}{
		"status": "error",
		"error":  errors,
	})
}
