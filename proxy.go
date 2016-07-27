package bardo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func (db *Database) Query(
	query string,
	args ...interface{},
) (*sql.Rows, error) {
	return db.Handle().Query(query, args...)
}

func (db *Database) Queryx(
	query string,
	args ...interface{},
) (*sqlx.Rows, error) {
	if db.transaction == nil {
		return db.db.Queryx(query, args...)
	}

	return db.transaction.Queryx(query, args...)
}

func (db *Database) QueryRowx(
	query string,
	args ...interface{},
) *sqlx.Row {
	if db.transaction == nil {
		return db.db.QueryRowx(query, args...)
	}

	return db.transaction.QueryRowx(query, args...)
}

func (db *Database) Exec(
	query string,
	args ...interface{},
) (sql.Result, error) {
	return db.Handle().Exec(query, args...)
}

func (db *Database) Get(
	dest interface{}, query string, args ...interface{},
) error {
	if db.transaction == nil {
		return db.db.Get(dest, query, args...)
	}

	return db.transaction.Get(dest, query, args...)
}

func (db *Database) Select(
	dest interface{}, query string, args ...interface{},
) error {
	if db.transaction == nil {
		return db.db.Select(dest, query, args...)
	}

	return db.transaction.Select(dest, query, args...)
}
