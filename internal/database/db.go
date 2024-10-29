package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"context"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func SetupDatabase() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
			log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Parse the connection string to add additional parameters
	connConfig, err := pgx.ParseConfig(dbURL)
	if err != nil {
			log.Fatalf("Failed to parse database URL: %v", err)
	}

	// Configure connection pool settings
	db, err := sql.Open("pgx", connConfig.ConnString())
	if err != nil {
			log.Fatalf("Failed to create database connection pool: %v", err)
	}

	// Set pool configuration
	db.SetMaxOpenConns(25)                  // Limit max open connections
	db.SetMaxIdleConns(10)                  // Maintain some idle connections
	db.SetConnMaxLifetime(30 * time.Minute) // Recycle connections periodically
	db.SetConnMaxIdleTime(5 * time.Minute)  // Close idle connections after time

	// Implement connection retry logic with backoff
	var lastErr error
	for attempts := 1; attempts <= 3; attempts++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			err := db.PingContext(ctx)
			cancel()
			
			if err == nil {
					fmt.Println("Successfully connected to database")
					return db
			}
			
			lastErr = err
			retryDelay := time.Duration(attempts) * 2 * time.Second
			fmt.Printf("Failed to connect to database (attempt %d): %v. Retrying in %v...\n", 
								attempts, err, retryDelay)
			time.Sleep(retryDelay)
	}

	log.Fatalf("Failed to connect to database after 3 attempts. Last error: %v", lastErr)
	return nil
}
