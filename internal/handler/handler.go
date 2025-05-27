package handler

import (
	"TaskManager/internal/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

// Package handler provides HTTP request handlers for the application.
type Handler struct {
	services *service.Service
	logger   *zap.Logger
}

// NewHandler creates a new Handler instance with the provided services.
func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// InitRoutes initializes the routes for the HTTP server.
func (h *Handler) InitRoutes(logger *zap.Logger) *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	logger.Info("Routes initialized")
	return e
}
