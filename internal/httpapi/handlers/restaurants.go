package handlers

import (
	"net/http"
	"strconv"

	reqdto "github.com/example/hungry-hub/internal/httpapi/handlers/request"
	"github.com/example/hungry-hub/internal/repository"
	"github.com/labstack/echo/v4"
)

type RestaurantsHandler struct {
	restaurants *repository.RestaurantRepo
	menuItems   *repository.MenuItemRepo
}

func NewRestaurants(restaurants *repository.RestaurantRepo, menuItems *repository.MenuItemRepo) *RestaurantsHandler {
	return &RestaurantsHandler{restaurants: restaurants, menuItems: menuItems}
}

func (h *RestaurantsHandler) Create(c echo.Context) error {
	var req reqdto.CreateRestaurantDTO
	if err := c.Bind(&req); err != nil {
		return badRequest(c, "invalid json body", nil)
	}

	details := map[string]any{}
	if req.Name == "" {
		details["name"] = "is required"
	}

	if req.Address == "" {
		details["address"] = "is required"
	}

	if len(details) > 0 {
		return badRequest(c, "validation failed", details)
	}

	out, err := h.restaurants.Create(c.Request().Context(), req.Name, req.Address, req.Phone, req.OpeningHours)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, out)
}

func (h *RestaurantsHandler) List(c echo.Context) error {
	out, err := h.restaurants.List(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, out)
}

func (h *RestaurantsHandler) Get(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return badRequest(c, "invalid id", nil)
	}

	rest, err := h.restaurants.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	items, err := h.menuItems.ListByRestaurant(c.Request().Context(), id, nil)
	if err != nil {
		return err
	}

	rest.MenuItems = items

	return c.JSON(http.StatusOK, rest)
}

func (h *RestaurantsHandler) Update(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return badRequest(c, "invalid id", nil)
	}

	var req reqdto.UpdateRestaurantDTO
	if err := c.Bind(&req); err != nil {
		return badRequest(c, "invalid json body", nil)
	}

	out, err := h.restaurants.Update(c.Request().Context(), id, req.Name, req.Address, req.Phone, req.OpeningHours)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, out)
}

func (h *RestaurantsHandler) Delete(c echo.Context) error {
	id, err := parseIDParam(c, "id")
	if err != nil {
		return badRequest(c, "invalid id", nil)
	}

	if err := h.restaurants.Delete(c.Request().Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func parseIDParam(c echo.Context, key string) (int64, error) {
	raw := c.Param(key)
	return strconv.ParseInt(raw, 10, 64)
}
