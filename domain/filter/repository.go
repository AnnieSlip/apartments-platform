package filter

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

type Repository interface {
	SaveFilter(ctx context.Context, userID int, filter models.ApartmentFilter) error
	GetFiltersByUser(ctx context.Context, userID int) ([]models.ApartmentFilter, error)
	GetAllFilters(ctx context.Context) ([]models.UserFilter, error)
	DeleteFilter(ctx context.Context, userID int) error
}

// EsRepository handles Elasticsearch storage
type EsRepository interface {
	SaveFilter(ctx context.Context, userID string, filter map[string]interface{}) error
}
