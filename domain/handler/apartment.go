package handler

import (
	"net/http"

	"github.com/ani-javakhishvili/apartments-platform/domain/apartment"
	"github.com/labstack/echo/v4"
)

type ApartmentHandler struct {
	Service *apartment.Service
}

func NewApartmentHandler(s *apartment.Service) *ApartmentHandler {
	return &ApartmentHandler{Service: s}
}

func (h *ApartmentHandler) ListApartments(c echo.Context) error {
	apts, err := h.Service.ListApartments(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, apts)
}

func (h *ApartmentHandler) CreateApartment(c echo.Context) error {
	var req apartment.Apartment
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	a, err := h.Service.CreateApartment(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, a)
}
