package db

import (
	"context"
	"database/sql"
	"strings" // Added import
	"time"

	_ "github.com/lib/pq" // Import the postgres driver
)

func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	// Ensure sslmode is disabled, overriding any potential environment issues for diagnosis
	connStr := addr
	if !strings.Contains(connStr, "sslmode=") {
		if strings.Contains(connStr, "?") {
			connStr += "&sslmode=disable" // Use & if other params exist
		} else {
			connStr += "?sslmode=disable"
		}
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		// It's helpful to close the db if duration parsing fails
		db.Close()
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		// Close the db connection if ping fails
		db.Close()
		return nil, err
	}

	return db, nil
}
