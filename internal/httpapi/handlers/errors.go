package handlers

import (
	"errors"
	"net/http"

	"github.com/example/hungry-hub/internal/repository"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}

func ErrorHandler(err error, c echo.Context) {
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code := he.Code
		msg := "request failed"
		if s, ok := he.Message.(string); ok && s != "" {
			msg = s
		}
		_ = c.JSON(code, ErrorResponse{Message: msg})

		return
	}

	if errors.Is(err, repository.ErrNotFound) {
		_ = c.JSON(http.StatusNotFound, ErrorResponse{Message: "not found"})
		return
	}

	_ = c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "internal server error"})
}

func badRequest(c echo.Context, message string, details map[string]any) error {
	return c.JSON(http.StatusBadRequest, ErrorResponse{Message: message, Details: details})
}
