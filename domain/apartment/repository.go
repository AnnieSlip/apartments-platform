package apartment

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/models"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Apartment, error)
	Create(ctx context.Context, a Apartment) (Apartment, error)
	GetApartmentsByFilter(ctx context.Context, filter models.ApartmentFilter) ([]Apartment, error)
}
