package bardo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func (db *Database) Query(
	query string,
	args ...interface{},
) (*sql.Rows, error) {
	return db.Handlex().Query(query, args...)
}

func (db *Database) Queryx(
	query string,
	args ...interface{},
) (*sqlx.Rows, error) {
	return db.Handlex().Queryx(query, args...)
}

func (db *Database) QueryRowx(
	query string,
	args ...interface{},
) *sqlx.Row {
	return db.Handlex().QueryRowx(query, args...)
}

func (db *Database) Exec(
	query string,
	args ...interface{},
) (sql.Result, error) {
	return db.Handlex().Exec(query, args...)
}

func (db *Database) Get(
	dest interface{}, query string, args ...interface{},
) error {
	return db.Handlex().Get(dest, query, args...)
}

func (db *Database) Select(
	dest interface{}, query string, args ...interface{},
) error {
	return db.Handlex().Select(dest, query, args...)
}
