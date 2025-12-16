package handler

import (
	"net/http"
	"strconv"

	"github.com/ani-javakhishvili/apartments-platform/domain/filter"
	"github.com/ani-javakhishvili/apartments-platform/domain/models"

	"github.com/labstack/echo/v4"
)

type FilterHandler struct {
	service *filter.Service
}

func NewFilterHandler(s *filter.Service) *FilterHandler {
	return &FilterHandler{service: s}
}

// CreateOrUpdateFilter handles POST /filters
func (h *FilterHandler) CreateOrUpdateFilter(c echo.Context) error {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid userID"})
	}

	var req models.ApartmentFilter
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.service.CreateOrUpdateFilter(c.Request().Context(), userID, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

// GetUserFilters handles GET /filters/:userID
func (h *FilterHandler) GetUserFilters(c echo.Context) error {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid userID"})
	}

	filters, err := h.service.GetFiltersByUser(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, filters)
}
