package apartment

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
	"github.com/labstack/gommon/log"
)

type Service struct {
	repo   Repository
	esRepo EsRepository
}

func NewService(r Repository, er EsRepository) *Service {
	return &Service{repo: r, esRepo: er}
}

func (s *Service) ListApartments(ctx context.Context) ([]models.Apartment, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) CreateApartment(ctx context.Context, a models.Apartment) (models.Apartment, error) {
	// todo: we can add validation here, if sth is emty in slice and so on...
	newA, err := s.repo.Create(ctx, a)
	if err != nil {
		return models.Apartment{}, err
	}

	go s.handleApartmentIndexed(newA)

	return newA, nil
}

func (s *Service) handleApartmentIndexed(a models.Apartment) {
	ctx := context.Background()

	// index apartment
	if err := s.esRepo.IndexApartment(ctx, a); err != nil {
		log.Errorf("ES index failed for apartment %d: %v", a.ID, err)
		// todo: retry / dead-letter queue / cron reindex
		// save this apartment ID in a Redis queue or Kafka topic for retry later
		// example:
		//   redis.LPush("apartment_index_retry", a.ID)
		//   or
		//   kafka.Produce("apartment_index_retry", a.ID)
		return
	}

	//  percolate apartment against filters
	userIDs, err := s.esRepo.PercolateApartment(ctx, a)
	if err != nil {
		log.Errorf("Percolation failed for apartment %d: %v", a.ID, err)
		return
	}

	// todo: notify users (optional)
	if len(userIDs) > 0 {
		// enqueue notifications
		log.Infof("Apartment %d matched users: %v", a.ID, userIDs)
	}
}
