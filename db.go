package bardo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

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

type Handle interface {
	Execer
	Queryer
}

type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// Handle returns the database/sql handle
func (db *Database) Handle() Handle {
	if db.transaction != nil {
		return db.transaction.Tx
	}

	return db.db.DB
}

// UnWrap returns the jmoiron/sqlx handle
func (db *Database) UnWrap() *sqlx.DB {
	return db.db
}
