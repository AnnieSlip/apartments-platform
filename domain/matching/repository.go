package matching

import "context"

type Repository interface {
	SaveUserMatches(ctx context.Context, userID int, apartmentIDs []int, week int) error
	GetUserMatches(ctx context.Context, userID int, week int) ([]int, error)
}
