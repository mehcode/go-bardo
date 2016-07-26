package bardo

import "github.com/jmoiron/sqlx"

type Database struct {
	db *sqlx.DB

	// Active transaction (nil if not inside a transaction)
	transaction *sqlx.Tx
}

// Wrap an existing sqlx DB
func Wrap(db *sqlx.DB) *Database {
	return &Database{
		db:          db.Unsafe(),
		transaction: nil,
	}
}

// Handle returns the database/sql handle
func (db *Database) Handle() interface{} {
	if db.transaction != nil {
		return db.transaction.Tx
	}

	return db.db.DB
}

// Handlex returns the jmoiron/sqlx handle
func (db *Database) Handlex() *sqlx.DB {
	return db.db
}
