package matching

import (
	"context"

	"github.com/ani-javakhishvili/apartments-platform/domain/storage/cassandra"
)

type CassandraRepo struct{}

func NewCassandraRepo() *CassandraRepo {
	return &CassandraRepo{}
}

func (r *CassandraRepo) SaveUserMatches(ctx context.Context, userID int, apartmentIDs []int, week int) error {
	query := `INSERT INTO user_matches (user_id, week, apartment_ids) VALUES (?, ?, ?)`
	return cassandra.Session.Query(query, userID, week, apartmentIDs).WithContext(ctx).Exec()
}
