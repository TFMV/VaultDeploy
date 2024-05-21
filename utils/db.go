package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDatabase initializes a PostgreSQL connection pool
func ConnectDatabase(config *Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.DBName,
		config.Postgres.SSLMode,
	)

	// Parse configuration
	pgxConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Unable to parse database configuration: %v\n", err)
		return nil, err
	}

	// Initialize connection pool
	dbpool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	return dbpool, nil
}
