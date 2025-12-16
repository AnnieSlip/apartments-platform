package handler

import (
	"net/http"

	"github.com/ani-javakhishvili/apartments-platform/domain/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *user.Service
}

// Initialize new handler
func NewUserHandler(s *user.Service) *UserHandler {
	return &UserHandler{Service: s}
}

// GET /users
func (h *UserHandler) ListUsers(c echo.Context) error {
	users, err := h.Service.ListUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

// POST /users
func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req user.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	u, err := h.Service.RegisterUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}
