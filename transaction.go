package bardo

import (
  "errors"
)

func (db *Database) Begin() error {
	if db.transaction != nil {
		return errors.New("transaction in progress")
	}

	var err error
	db.transaction, err = db.db.Beginx()
	if err != nil {
		return err
	}

	// NOTE: Unsafe just turns off fatal panics if our database
	//       returns more columns than we can map. This not unsafe at all and
	//       is just sqlx being opinionated (json unmarshal for instance is like
	//       this by default)
	db.transaction = db.transaction.Unsafe()

	return nil
}

func (db *Database) Commit() error {
	if db.transaction == nil {
		return errors.New("outside transaction")
	}

	err := db.transaction.Commit()
	if err != nil {
		return err
	}

	// Wipe transaction and move back to the DB
	db.transaction = nil

	return nil
}

func (db *Database) Rollback() error {
	if db.transaction == nil {
		return errors.New("outside transaction")
	}

	err := db.transaction.Rollback()
	if err != nil {
		return err
	}

	// Wipe transaction and move back to the DB
	db.transaction = nil

	return nil
}

func (db *Database) InTransaction() bool {
	return db.transaction != nil
}
