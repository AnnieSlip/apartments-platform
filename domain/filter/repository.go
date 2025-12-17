package filter

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

type Repository interface {
	SaveFilter(ctx context.Context, userID int, filter models.ApartmentFilter) error
	GetFiltersByUser(ctx context.Context, userID int) ([]models.ApartmentFilter, error)
	GetAllFilters(ctx context.Context) ([]models.UserFilter, error)
}

type MatchingRepository interface {
	SaveUserMatches(ctx context.Context, userID int, apartmentIDs []int, week int) error
}
