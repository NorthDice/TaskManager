package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// register and login handlers for authentication
func (h *Handler) register(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{
		"id": "12345",
	})
}

// login handler for authentication
func (h *Handler) login(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]interface{}{
		"Yes": "You are logged in",
	})
}
