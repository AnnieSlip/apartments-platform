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

	// save to Elasticsearch via injected repository
	if err := s.esRepo.SaveFilter(ctx, userIDStr, f); err != nil {

		return fmt.Errorf("failed to save filter in Elasticsearch: %w", err)
	}

	return nil
}
