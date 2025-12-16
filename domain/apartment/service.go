package apartment

import "context"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) ListApartments(ctx context.Context) ([]Apartment, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) CreateApartment(ctx context.Context, a Apartment) (Apartment, error) {
	// you can add validation or business rules here
	return s.repo.Create(ctx, a)
}

func (s *Service) GetByFilter(ctx context.Context, a Apartment) (Apartment, error) {
	// you can add validation or business rules here
	return s.repo.Create(ctx, a)
}
