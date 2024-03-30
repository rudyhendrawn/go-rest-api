package database

import (
	"database/sql"
	"fmt"
	"go-rest-api/internal/pkg/config"

	_ "github.com/lib/pq"
)

// ConnectDB initializes a connection to the database
func ConnectDB(cfg *config.AppConfig) (*sql.DB, error) {
	// Construct the connection string
	var dbName, dbHost, dbPort, dbUser, dbPassword string
	dbHost = cfg.DBHost
	dbName = cfg.DBName
	dbPort = cfg.DBPort
	dbUser = cfg.DBUser
	dbPassword = cfg.DBPassword

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
