package main

import (
	"context"
	"log"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
	esStorage "github.com/ani-javakhishvili/apartments-platform/domain/storage/elasticsearch"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"

	"github.com/ani-javakhishvili/apartments-platform/domain/storage/postgres"

	apartmentRepo "github.com/ani-javakhishvili/apartments-platform/domain/apartment"
)

type WeeklyJob struct {
	apartmentRepo apartmentRepo.Repository
	esRepo        EsRepository
}

type EsRepository interface {
	PercolateApartment(ctx context.Context, a models.Apartment) ([]string, error)
}

func NewWeeklyJob(
	apartmentRepo apartmentRepo.Repository,
	esRepo EsRepository,

) *WeeklyJob {
	return &WeeklyJob{
		apartmentRepo: apartmentRepo,
		esRepo:        esRepo,
	}
}

func main() {
	ctx := context.Background()
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	})
	if err != nil {
		log.Fatalf("elasticsearch client init failed: %v", err)
	}

	aptRepo := postgres.NewApartmentPostgresRepo()
	esRepo := esStorage.NewEsRepo(esClient)

	job := NewWeeklyJob(aptRepo, esRepo)

	if err := job.Run(ctx); err != nil {
		log.Fatalf("Weekly job failed: %v", err)
	}

}

func (j *WeeklyJob) Run(ctx context.Context) error {

	apartments, err := j.apartmentRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	userMatches := map[string][]int{} // userID → apartmentIDs

	for _, a := range apartments {
		users, err := j.esRepo.PercolateApartment(ctx, a)
		if err != nil {
			log.Printf("percolation failed for apartment %d: %v", a.ID, err)
			continue
		}

		for _, userID := range users {
			userMatches[userID] = append(userMatches[userID], a.ID)
		}
	}

	for userID, apartmentIDs := range userMatches {
		sendNotification(userID, apartmentIDs)
	}

	return nil
}

func sendNotification(userID string, apartmentIDs []int) {
	// here goes email sending logic -  with email, push, etc.
	log.Printf("Sending notification to user %d: matched apartments %v\n", userID, apartmentIDs)
}
