package user

import "context"

type Repository interface {
	GetAll(ctx context.Context) (Users, error)
	Create(ctx context.Context, email string) (*User, error)
}
