package filter

import (
	"context"
	"time"

	"github.com/ani-javakhishvili/apartments-platform/domain/apartment"
	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

type Service struct {
	repo    Repository
	aptRepo apartment.Repository
}

func NewService(repo Repository, aptRepo apartment.Repository) *Service {
	return &Service{repo: repo, aptRepo: aptRepo}
}

func (s *Service) GetFiltersByUser(ctx context.Context, userID int) ([]models.ApartmentFilter, error) {
	return s.repo.GetFiltersByUser(ctx, userID)
}

func (s *Service) CreateOrUpdateFilter(ctx context.Context, userID int, filter models.ApartmentFilter) error {
	// save user filter in db
	if err := s.repo.SaveFilter(ctx, userID, filter); err != nil {
		return err
	}
	return nil
}

func currentWeek() int {
	_, week := time.Now().ISOWeek()
	return week
}
