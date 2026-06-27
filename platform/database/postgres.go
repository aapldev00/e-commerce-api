// Package database provides tools for external data store connectivity.
package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresPool initializes a new PostgreSQL connection pool using pgx.
// It parses the connection string and verifies the connectivity with a ping.
func NewPostgresPool(connString string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Parse the connection string into a configuration object.
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}

	// Set defaults for the connection pool.
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 5 * time.Minute

	// Create the pool with the specific configuration.
	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Verify the database is reachable.
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("database is unreachable: %w", err)
	}

	return db, nil
}
