package main

import (
	"context"
	"log"

	"github.com/ani-javakhishvili/apartments-platform/domain/filter"
	"github.com/ani-javakhishvili/apartments-platform/domain/storage/cassandra"
	"github.com/ani-javakhishvili/apartments-platform/domain/storage/postgres"
)

func main() {
	ctx := context.Background()

	// connect dbs
	if err := postgres.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	if err := cassandra.Connect(); err != nil {
		log.Fatal(err)
	}

	filterRepo := postgres.NewFilterPostgresRepo()
	aptRepo := postgres.NewApartmentPostgresRepo()
	matchRepo := cassandra.NewRepository(cassandra.Session)

	service := filter.NewService(filterRepo, aptRepo, matchRepo)

	log.Println("Precompute job started")

	if err := service.RecomputeAll(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Precompute job finished")
}
