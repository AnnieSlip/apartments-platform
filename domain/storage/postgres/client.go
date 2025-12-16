package postgres

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect(ctx context.Context) error {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/apartments?sslmode=disable"
	}

	var err error
	DB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}

	return DB.Ping(ctx)
}
