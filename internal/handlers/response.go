package handlers

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c echo.Context, logger *zap.Logger, statusCode int, message string) {
	logger.Error("Request failed", zap.String("message", message))
	c.JSON(statusCode, errorResponse{Message: message})
}
