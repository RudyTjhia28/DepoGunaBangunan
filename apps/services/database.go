package services

import (
	"database/sql"
	"depogunabangunan/config"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDatabase establishes a connection to the PostgreSQL database
func ConnectDatabase(cfg *config.Configuration) (*sql.DB, error) {
	// Build the PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSchema)

	// Establish the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Set maximum open and idle connections
	db.SetMaxOpenConns(cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	fmt.Println("Connected to the PostgreSQL database")

	return db, nil
}
