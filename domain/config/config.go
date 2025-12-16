package config

import (
	"log"

	"github.com/ani-javakhishvili/apartments-platform/domain/apartment"
	"github.com/ani-javakhishvili/apartments-platform/domain/handler"
	"github.com/ani-javakhishvili/apartments-platform/domain/storage/postgres"
	"github.com/ani-javakhishvili/apartments-platform/domain/user"
)

// App holds all services and handlers
type App struct {
	UserHandler      *handler.UserHandler
	ApartmentHandler *handler.ApartmentHandler
}

// Initialize connects to DB and wires repositories, services, and handlers
func Init() *App {
	// Connect to Postgres
	if err := postgres.Connect(); err != nil {
		log.Fatalf("Postgres connection failed: %v", err)
	}

	// Repositories
	userRepo := postgres.NewUserPostgresRepo()
	aptRepo := postgres.NewApartmentPostgresRepo()

	// Services
	userService := user.NewService(userRepo)
	aptService := apartment.NewService(aptRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	aptHandler := handler.NewApartmentHandler(aptService)

	return &App{
		UserHandler:      userHandler,
		ApartmentHandler: aptHandler,
	}
}
