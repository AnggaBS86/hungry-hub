package handlers

import (
	"net/http"

	reqdto "github.com/example/hungry-hub/internal/httpapi/handlers/request"
	"github.com/example/hungry-hub/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type MenuItemsHandler struct {
	restaurants *repository.RestaurantRepo
	menuItems   *repository.MenuItemRepo
}

func NewMenuItems(restaurants *repository.RestaurantRepo, menuItems *repository.MenuItemRepo) *MenuItemsHandler {
	return &MenuItemsHandler{restaurants: restaurants, menuItems: menuItems}
}

func (h *MenuItemsHandler) CreateForRestaurant(c echo.Context) error {
	restaurantID, err := parseIDParam(c, "id")
	if err != nil {
		return badRequest(c, "invalid restaurant id", nil)
	}

	if _, err := h.restaurants.Get(c.Request().Context(), restaurantID); err != nil {
		return err
	}

	var req reqdto.CreateMenuItemDTO
	if err := c.Bind(&req); err != nil {
		return badRequest(c, "invalid json body", nil)
	}

	details := map[string]any{}
	if req.Name == "" {
		details["name"] = "is required"
	}

	if req.Price == "" {
		details["price"] = "is required"
	}

	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		details["price"] = "must be a decimal string"
	} else if price.LessThanOrEqual(decimal.Zero) {
		details["price"] = "must be > 0"
	}

	if len(details) > 0 {
		return badRequest(c, "validation failed", details)
	}

	out, err := h.menuItems.Create(c.Request().Context(), restaurantID, req.Name, req.Description, price, req.Category, req.IsAvailable)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, out)
}

func (h *MenuItemsHandler) ListForRestaurant(c echo.Context) error {
	restaurantID, err := parseIDParam(c, "id")
	if err != nil {
		return badRequest(c, "invalid restaurant id", nil)
	}

	if _, err := h.restaurants.Get(c.Request().Context(), restaurantID); err != nil {
		return err
	}

	var category *string
	if v := c.QueryParam("category"); v != "" {
		category = &v
	}

	out, err := h.menuItems.ListByRestaurant(c.Request().Context(), restaurantID, category)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, out)
}

func (h *MenuItemsHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return badRequest(c, "invalid menu item id", nil)
	}

	var req reqdto.UpdateMenuItemDTO
	if err := c.Bind(&req); err != nil {
		return badRequest(c, "invalid json body", nil)
	}

	var price *decimal.Decimal
	if req.Price != nil {
		p, err := decimal.NewFromString(*req.Price)
		if err != nil {
			return badRequest(c, "validation failed", map[string]any{"price": "must be a decimal string"})
		}
		price = &p
	}

	out, err := h.menuItems.Update(c.Request().Context(), id, req.Name, req.Description, price, req.Category, req.IsAvailable)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, out)
}

func (h *MenuItemsHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return badRequest(c, "invalid menu item id", nil)
	}

	if err := h.menuItems.Delete(c.Request().Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
