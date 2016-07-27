package bardo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// CreateDatabase ...
// Creates the database on the SQL server (if not exists)
func CreateDatabase(url string) error {
	// Find database name (that we will create)
	dbname := GetDBNameFromURL(url)

	// Get a connection string to the template DB
	templateURL := ReplaceDBNameInURL(url, "postgres")

	// Connect to the new template DB
	db, err := sqlx.Connect("postgres", templateURL)
	if err != nil {
		return err
	}

	// Check if the database exists
	var exists bool
	err = db.QueryRow(`SELECT TRUE FROM pg_database WHERE datname = $1`, dbname).
		Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if exists {
		// Don't continue if it already exists
		return nil
	}

	// Create the database
	_, err = db.Exec(`CREATE DATABASE "` + dbname + `"`)
	if err != nil {
		return err
	}

	// Close the connection pool
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}

// DropDatabase ...
func DropDatabase(url string) error {
	// Find database name (that we will create)
	dbname := GetDBNameFromURL(url)

	// Get a connection string to the template DB
	templateURL := ReplaceDBNameInURL(url, "postgres")

	// Connect to the new template DB
	db, err := sqlx.Connect("postgres", templateURL)
	if err != nil {
		return err
	}

	// Terminate all connections (abrubtly) for this database
	_, err = db.Exec(`
		SELECT pg_terminate_backend(pg_stat_activity.pid)
		FROM pg_stat_activity
		WHERE pg_stat_activity.datname = $1
			AND pid <> pg_backend_pid()
	`, dbname)
	if err != nil {
		return err
	}

	// Drop the database
	_, err = db.Exec(`DROP DATABASE IF EXISTS "` + dbname + `"`)
	if err != nil {
		return err
	}

	// Close the connection pool
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
