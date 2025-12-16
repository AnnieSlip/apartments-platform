package user

import "context"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) ListUsers(c context.Context) ([]User, error) {
	// could add business rules here if needed
	return s.repo.GetAll(c)
}

func (s *Service) RegisterUser(c context.Context, req User) (*User, error) {
	// could add validation: check email format or uniqueness before saving
	res, err := s.repo.Create(c, req.Email)
	if err != nil {
		return nil, err
	}
	return res, nil
}
