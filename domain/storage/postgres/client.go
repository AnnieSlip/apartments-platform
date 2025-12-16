package postgres

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	// Read env variables
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "app"
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "app"
	}
	dbname := os.Getenv("POSTGRES_DB")
	if dbname == "" {
		dbname = "apartments"
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	// Connect with a pool
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("unable to connect to postgres: %w", err)
	}

	// Ping to check connection
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("unable to ping postgres: %w", err)
	}

	DB = pool
	fmt.Println("Connected to PostgreSQL successfully")
	return nil
}
