package filter

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

type Repository interface {
	SaveFilter(ctx context.Context, userID int, filter models.ApartmentFilter) error
	GetFilter(ctx context.Context, userID int) (ApartmentFilter, error)
}
