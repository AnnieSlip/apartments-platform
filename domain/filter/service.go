package filter

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ani-javakhishvili/apartments-platform/domain/apartment"
	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

type Service struct {
	repo    Repository
	esRepo  EsRepository
	aptRepo apartment.Repository
}

// NewService creates a new filter service with injected dependencies
func NewService(repo Repository, aptRepo apartment.Repository, esRepo EsRepository) *Service {
	return &Service{
		repo:    repo,
		aptRepo: aptRepo,
		esRepo:  esRepo,
	}
}

// GetFiltersByUser retrieves user filters from Postgres
func (s *Service) GetFiltersByUser(ctx context.Context, userID int) ([]models.ApartmentFilter, error) {
	return s.repo.GetFiltersByUser(ctx, userID)
}

// CreateOrUpdateFilter saves a user filter to Postgres and Elasticsearch
func (s *Service) CreateOrUpdateFilter(ctx context.Context, userID int, f models.ApartmentFilter) error {
	userIDStr := strconv.Itoa(userID)

	// convert filter struct to ES percolator query
	esFilter := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": []interface{}{
				map[string]interface{}{"term": map[string]interface{}{"city": f.City}},
				map[string]interface{}{"range": map[string]interface{}{"price": map[string]interface{}{"lte": f.MaxPrice}}},
				map[string]interface{}{"term": map[string]interface{}{"rooms": f.RoomNumbers}},
				map[string]interface{}{"term": map[string]interface{}{"bedrooms": f.BedroomNumbers}},
				map[string]interface{}{"term": map[string]interface{}{"bathrooms": f.BathroomNumbers}},
			},
		},
	}

	// save to Postgres
	if err := s.repo.SaveFilter(ctx, userID, f); err != nil {
		return fmt.Errorf("failed to save filter in Postgres: %w", err)
	}

	// save to Elasticsearch via injected repository
	if err := s.esRepo.SaveFilter(ctx, userIDStr, esFilter); err != nil {
		// TODO: think about improvement here
		// rollback Postgres if ES fails
		if rollbackErr := s.repo.DeleteFilter(ctx, userID); rollbackErr != nil {
			return fmt.Errorf("failed to save filter in ES: %v, also failed to rollback Postgres: %v", err, rollbackErr)
		}
		return fmt.Errorf("failed to save filter in Elasticsearch: %w", err)
	}

	return nil
}
