package config

import (
	"context"
	"log"

	"github.com/ani-javakhishvili/apartments-platform/domain/apartment"
	"github.com/ani-javakhishvili/apartments-platform/domain/filter"
	"github.com/ani-javakhishvili/apartments-platform/domain/handler"
	esStorage "github.com/ani-javakhishvili/apartments-platform/domain/storage/elasticsearch"
	"github.com/ani-javakhishvili/apartments-platform/domain/storage/postgres"
	"github.com/ani-javakhishvili/apartments-platform/domain/user"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

type App struct {
	UserHandler      *handler.UserHandler
	ApartmentHandler *handler.ApartmentHandler
	FilterHandler    *handler.FilterHandler
	FilterService    *filter.Service
	ESClient         *elasticsearch.Client
}

// initialize connects to DB and wires repositories, services, and handlers
func Init() *App {
	ctx := context.Background()
	if err := postgres.Connect(ctx); err != nil {
		log.Fatalf("postgres connection failed: %v", err)
	}

	// elastic search
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	})
	if err != nil {
		log.Fatalf("elasticsearch client init failed: %v", err)
	}

	// create indices if they don't exist
	if err := esStorage.CreateIndices(esClient); err != nil {
		log.Fatalf("failed to create ES indices: %v", err)
	}
	// repositories
	userRepo := postgres.NewUserPostgresRepo()
	aptRepo := postgres.NewApartmentPostgresRepo()
	filterRepo := postgres.NewFilterPostgresRepo()
	esRepo := esStorage.NewEsRepo(esClient)

	// services
	userService := user.NewService(userRepo)
	aptService := apartment.NewService(aptRepo)
	filterService := filter.NewService(filterRepo, aptRepo, esRepo)

	// handlers
	userHandler := handler.NewUserHandler(userService)
	aptHandler := handler.NewApartmentHandler(aptService)
	filterHandler := handler.NewFilterHandler(filterService)

	return &App{
		UserHandler:      userHandler,
		ApartmentHandler: aptHandler,
		FilterHandler:    filterHandler,
		FilterService:    filterService,
		ESClient:         esClient,
	}
}
