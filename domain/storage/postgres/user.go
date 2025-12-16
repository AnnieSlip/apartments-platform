package postgres

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/user"
)

type UserDB struct {
	ID    int
	Email string
}

type UsersDB []UserDB

// UserPostgresRepo is the repository
type UserPostgresRepo struct{}

// NewUserPostgresRepo creates a new repo instance
func NewUserPostgresRepo() *UserPostgresRepo {
	return &UserPostgresRepo{}
}

// GetAll fetches all users
func (r *UserPostgresRepo) GetAll(ctx context.Context) (user.Users, error) {
	rows, err := DB.Query(ctx, "SELECT id, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users UsersDB
	for rows.Next() {
		var u UserDB
		if err := rows.Scan(&u.ID, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return mapUsersToDomain(users), nil
}

// Create inserts a new user
func (r *UserPostgresRepo) Create(ctx context.Context, email string) (*user.User, error) {
	var u UserDB
	err := DB.QueryRow(ctx,
		"INSERT INTO users(email) VALUES($1) RETURNING id, email",
		email,
	).Scan(&u.ID, &u.Email)

	if err != nil {
		return nil, err
	}
	res := mapUserToDomain(u)
	return &res, nil
}

func mapUserToDomain(u UserDB) user.User {
	return user.User{
		ID:    u.ID,
		Email: u.Email,
	}
}

func mapUsersToDomain(us UsersDB) user.Users {
	if len(us) == 0 {
		return user.Users{}
	}
	var res user.Users
	for _, u := range us {
		res = append(res, mapUserToDomain(u))
	}
	return res
}
