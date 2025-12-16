package main

import (
	"log"
	"net/http"

	"github.com/ani-javakhishvili/apartments-platform/domain/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := config.Init()

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// User routes
	e.GET("/users", app.UserHandler.ListUsers)
	e.POST("/users", app.UserHandler.RegisterUser)

	// Apartment routes
	e.GET("/apartments", app.ApartmentHandler.ListApartments)
	e.POST("/apartments", app.ApartmentHandler.CreateApartment)

	log.Println("API server running on :8080")
	log.Fatal(e.Start(":8080"))
}
