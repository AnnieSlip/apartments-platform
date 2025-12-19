package apartment

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Apartment, error)
	Create(ctx context.Context, a Apartment) (Apartment, error)
	// GetApartmentsByFilter(ctx context.Context, filter models.ApartmentFilter) ([]Apartment, error)
}
