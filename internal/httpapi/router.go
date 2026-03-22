package httpapi

import (
	"net/http"

	"github.com/example/hungry-hub/internal/httpapi/handlers"
	"github.com/example/hungry-hub/internal/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{"status": "ok"})
	})

	restaurantsRepo := repository.NewRestaurantRepo(db)
	menuItemsRepo := repository.NewMenuItemRepo(db)

	restaurants := handlers.NewRestaurants(restaurantsRepo, menuItemsRepo)
	menuItems := handlers.NewMenuItems(restaurantsRepo, menuItemsRepo)

	e.POST("/restaurants", restaurants.Create)
	e.GET("/restaurants", restaurants.List)
	e.GET("/restaurants/:id", restaurants.Get)
	e.PUT("/restaurants/:id", restaurants.Update)
	e.DELETE("/restaurants/:id", restaurants.Delete)

	e.POST("/restaurants/:id/menu_items", menuItems.CreateForRestaurant)
	e.GET("/restaurants/:id/menu_items", menuItems.ListForRestaurant)
	e.PUT("/menu_items/:id", menuItems.Update)
	e.DELETE("/menu_items/:id", menuItems.Delete)
}
