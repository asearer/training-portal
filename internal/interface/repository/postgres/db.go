package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// ConnectDB initializes the global DB connection using environment variables or config.
func ConnectDB() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("database environment variables are not set")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	db = conn
	return db, nil
}

// GetDB returns the global DB connection (must call ConnectDB first).
func GetDB() *sql.DB {
	return db
}
