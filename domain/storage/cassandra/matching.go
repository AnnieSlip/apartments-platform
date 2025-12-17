package cassandra

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
)

type MatchingCassandraRepo struct {
	session *gocql.Session
}

func NewRepository(session *gocql.Session) *MatchingCassandraRepo {
	return &MatchingCassandraRepo{session: session}
}

func (r *MatchingCassandraRepo) SaveUserMatches(ctx context.Context, userID int, apartmentIDs []int, week int) error {
	if err := r.session.Query(
		`INSERT INTO user_matches (user_id, week, apartment_ids) VALUES (?, ?, ?)`,
		userID, week, apartmentIDs,
	).WithContext(ctx).Exec(); err != nil {
		return fmt.Errorf("failed to save user matches: %w", err)
	}
	return nil
}

func (r *MatchingCassandraRepo) GetUserMatches(ctx context.Context, userID int, week int) ([]int, error) {
	var aptIDs []int
	if err := r.session.Query(
		`SELECT apartment_ids FROM user_matches WHERE user_id = ? AND week = ?`,
		userID, week,
	).WithContext(ctx).Scan(&aptIDs); err != nil {
		return nil, fmt.Errorf("failed to get user matches: %w", err)
	}
	return aptIDs, nil
}

func (r *MatchingCassandraRepo) GetAllMatchesForWeek(ctx context.Context, week int) (map[int][]int, error) {
	query := `SELECT user_id, apartment_ids FROM user_matches WHERE week = ?`
	iter := r.session.Query(query, week).WithContext(ctx).Iter()

	results := make(map[int][]int)
	var userID int
	var apartmentIDs []int

	for iter.Scan(&userID, &apartmentIDs) {
		results[userID] = apartmentIDs
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return results, nil
}
