package config

import (
	"context"
	"log"

	"github.com/ani-javakhishvili/apartments-platform/domain/apartment"
	"github.com/ani-javakhishvili/apartments-platform/domain/filter"
	"github.com/ani-javakhishvili/apartments-platform/domain/handler"
	"github.com/ani-javakhishvili/apartments-platform/domain/storage/cassandra"
	"github.com/ani-javakhishvili/apartments-platform/domain/storage/postgres"
	"github.com/ani-javakhishvili/apartments-platform/domain/user"
)

type App struct {
	UserHandler      *handler.UserHandler
	ApartmentHandler *handler.ApartmentHandler
	FilterHandler    *handler.FilterHandler
	FilterService    *filter.Service
}

// initialize connects to DB and wires repositories, services, and handlers
func Init() *App {
	// connect to Postgres
	ctx := context.Background()
	if err := postgres.Connect(ctx); err != nil {
		log.Fatalf("postgres connection failed: %v", err)
	}
	// initialize cassandra
	if err := cassandra.Connect(); err != nil {
		log.Fatalf("cassandra connection failed: %v", err)
	}
	// repositories
	userRepo := postgres.NewUserPostgresRepo()
	aptRepo := postgres.NewApartmentPostgresRepo()
	filterRepo := postgres.NewFilterPostgresRepo()
	matchRepo := cassandra.NewRepository(cassandra.Session)

	// services
	userService := user.NewService(userRepo)
	aptService := apartment.NewService(aptRepo)
	filterService := filter.NewService(filterRepo, aptRepo, matchRepo)

	// handlers
	userHandler := handler.NewUserHandler(userService)
	aptHandler := handler.NewApartmentHandler(aptService)
	filterHandler := handler.NewFilterHandler(filterService)

	return &App{
		UserHandler:      userHandler,
		ApartmentHandler: aptHandler,
		FilterHandler:    filterHandler,
		FilterService:    filterService,
	}
}
