package handlers

import (
	"TaskManager/internal/domain/model"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

// register and login handlers for authentication
func (h *Handler) register(e echo.Context) error {
	var input model.User
	log := h.logger.With(
		zap.String("handler", "register"),
	)

	if err := e.Bind(&input); err != nil {
		newErrorResponse(e, log, http.StatusBadRequest, err.Error())
		return nil
	}

	log.Info("Registering user", zap.String("username", input.Username))

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}

	log.Info("User registered successfully", zap.String("userID", string(id)))

	return e.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password"  bson:"password"`
}

// login handler for authentication
func (h *Handler) login(e echo.Context) error {
	var input signInInput

	log := h.logger.With(
		zap.String("handler", "login"),
	)

	if err := e.Bind(&input); err != nil {
		newErrorResponse(e, log, http.StatusBadRequest, err.Error())
		return nil
	}
	log.Info("User login attempt", zap.String("username", input.Username))

	token, err := h.services.Authorization.GenerateToken(
		input.Username,
		input.Password,
	)
	if err != nil {
		newErrorResponse(e, log, http.StatusInternalServerError, err.Error())
		return nil
	}

	log.Info("User logged in successfully")

	return e.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
