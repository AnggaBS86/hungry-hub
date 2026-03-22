package httpapi

import (
	"github.com/example/hungry-hub/internal/httpapi/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.HTTPErrorHandler = handlers.ErrorHandler

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.BodyLimit("2M"))

	RegisterRoutes(e, db)

	return e
}
