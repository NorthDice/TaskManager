package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	isEmptyString = ""
)
const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentityMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get(authorizationHeader)
		if header == isEmptyString {
			newErrorResponse(c, h.logger, http.StatusUnauthorized, "No authorization header provided")
			return nil
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(c, h.logger, http.StatusUnauthorized, "Invalid authorization header")
			return nil
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(c, h.logger, http.StatusUnauthorized, "Invalid authorization token")
			return nil
		}

		c.Set(userCtx, userId)
		return next(c)
	}
}
