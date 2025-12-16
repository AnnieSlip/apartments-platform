package matching

import "context"

type Repository interface {
	GetUserMatches(ctx context.Context, userID int, week int) ([]int, error)
	SaveUserMatches(ctx context.Context, userID int, apartmentIDs []int, week int) error
}
