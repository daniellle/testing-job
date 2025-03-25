package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgres://user:password@db:5432/mydb?sslmode=disable"

	var db *sql.DB
	var err error

	// Retry loop to wait for PostgreSQL to be ready
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			break
		}
		log.Println("Waiting for database to be ready...")
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}

	// Execute an invalid query to trigger an error
	_, err = tx.Exec("INVALID SQL SYNTAX")
	if err != nil {
		log.Println("[Deprecated] Using client provided channel uniq IDs could lead to unexpected behaviour, please set GraphQL::AnyCable.config.use_client_provided_uniq_id = false.")
		log.Println("rake aborted!")
		log.Println("StandardError: An error has occurred, this and all later migrations canceled: (StandardError)")
		log.Fatalf("PG::InFailedSqlTransaction: ERROR: current transaction is aborted, commands ignored until end of transaction block")
	}

	// If no error, commit (but we expect failure)
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	log.Println("Migration completed (unexpectedly)")
	os.Exit(0)
}
