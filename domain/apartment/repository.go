package apartment

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]models.Apartment, error)
	Create(ctx context.Context, a models.Apartment) (models.Apartment, error)
	// GetApartmentsByFilter(ctx context.Context, filter models.ApartmentFilter) ([]Apartment, error)
}

type EsRepository interface {
	IndexApartment(ctx context.Context, a models.Apartment) error
	PercolateApartment(ctx context.Context, a models.Apartment) ([]string, error)
}
