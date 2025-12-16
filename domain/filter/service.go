package filter

import (
	"context"
	"time"

	"github.com/ani-javakhishvili/apartments-platform/domain/apartment"
	"github.com/ani-javakhishvili/apartments-platform/domain/models"

	"github.com/ani-javakhishvili/apartments-platform/domain/matching"
)

type Service struct {
	repo      Repository
	aptRepo   apartment.Repository
	matchRepo matching.Repository
}

func NewService(repo Repository, aptRepo apartment.Repository, matchRepo matching.Repository) *Service {
	return &Service{repo: repo, aptRepo: aptRepo, matchRepo: matchRepo}
}

func (s *Service) CreateOrUpdateFilter(ctx context.Context, userID int, filter models.ApartmentFilter) error {
	// 1️⃣ save filter in Postgres
	if err := s.repo.SaveFilter(ctx, userID, filter); err != nil {
		return err
	}

	// 2️⃣ get matching apartments
	apartments, err := s.aptRepo.GetApartmentsByFilter(ctx, filter)
	if err != nil {
		return err
	}

	// 3️⃣ save matches in Cassandra
	aptIDs := make([]int, len(apartments))
	for i, a := range apartments {
		aptIDs[i] = a.ID
	}

	return s.matchRepo.SaveUserMatches(ctx, userID, aptIDs, currentWeek())
}

func currentWeek() int {
	_, week := time.Now().ISOWeek()
	return week
}
