package handlers

import (
	"TaskManager/internal/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Package handler provides HTTP request handlers for the application.
type Handler struct {
	services *service.Service
	logger   *zap.Logger
}

// NewHandler creates a new Handler instance with the provided services.
func NewHandler(services *service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

// InitRoutes initializes the routes for the HTTP server.
func (h *Handler) InitRoutes(logger *zap.Logger) *echo.Echo {
	e := echo.New()

	e.POST("/register", h.register)
	e.POST("/login", h.login)

	auth := e.Group("/tasks")
	auth.GET("", h.getTasks)
	auth.GET("/:id", h.getTaskByID)
	auth.POST("", h.createTask)
	auth.PUT("/:id", h.updateTask)
	auth.DELETE("/:id", h.deleteTask)

	logger.Info("Routes initialized")

	return e
}
