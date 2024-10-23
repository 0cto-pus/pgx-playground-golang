package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	ErrorDescription string `json:"errorDescription"`
}

func ErrorMessage(ctx echo.Context, status int, err error) error {
	return ctx.JSON(status, map[string]interface{}{
		"error": err.Error(),
	})
}

func InternalError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		"error": err.Error(),
	})
}

func BadRequestError(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": msg,
	})
}

func SuccessResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": "success",
		"data":    data,
	}
}