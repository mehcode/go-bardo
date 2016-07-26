package bardo

import (
	"strings"
)

// GetTables ...
// Get an array of all tables in the database
func (db *Database) getAllTables() ([]string, error) {
	var tables []struct {
		Name string `db:"table_name"`
	}

	err := db.Select(&tables, `
		SELECT table_name
		  FROM information_schema.tables
		 WHERE table_schema = 'public'
		   AND table_type = 'BASE TABLE'
			 AND table_name != 'goose_db_version'
			 AND table_name != 'client'
	`)

	if err != nil {
		return nil, err
	}

	var names []string
	for _, table := range tables {
		names = append(names, table.Name)
	}

	return names, nil
}

// Truncate ...
// Truncates all tables in the database
// TODO: Should accept exclusion list
func (db *Database) Truncate() error {
	names, err := db.getAllTables()
	if err != nil {
		return err
	}

	var tablenames []string
	for _, name := range names {
		tablenames = append(tablenames, `"`+name+`"`)
	}

	_, err = db.Exec(`
		TRUNCATE TABLE ` + strings.Join(tablenames, ", ") + `
		RESTART IDENTITY CASCADE
	`)

	if err != nil {
		return err
	}

	return nil
}
