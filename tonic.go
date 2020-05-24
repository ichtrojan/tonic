package tonic

import (
	"context"
	"database/sql"
	"sync"
)

// DB represents a database connection
type DB struct {
	sync.RWMutex
	Error error

	db     *sql.DB
	search *search
}

// Open initializes a new database connection
func Open(driver, source string) (db *DB, err error) {
	dbConn, err := sql.Open(driver, source)
	if err != nil {
		return
	}
	db = &DB{
		db: dbConn,
	}
	// Ping database connection
	if err = dbConn.Ping(); err != nil {
		dbConn.Close()
		return
	}
	return
}

// DB returns the underlying *sql.DB object in connection
func (d *DB) DB() *sql.DB {
	return d.db
}

// SetError set DB.Error err
func (d *DB) SetError(err error) error {
	if err != nil {
		d.Lock()
		d.Error = err
		d.Unlock()
	}
	return err
}

// Close closes current db connection.
func (d *DB) Close() error {
	return d.db.Close()
}

// Where return a new relation, filter records with given conditions
func (d *DB) Where(query interface{}, args ...interface{}) *DB {
	return d
}

// First finds the first record matching conditions
func (d *DB) First(out interface{}, args ...interface{}) *DB {
	return d
}

// Last finds the last record matching conditions
func (d *DB) Last(out interface{}, args ...interface{}) *DB {
	return d
}

// Find finds records matching conditions
func (d *DB) Find(out interface{}, args ...interface{}) *DB {
	return d
}

// Or filter records that match before conditions or current one
func (d *DB) Or(query interface{}, args ...interface{}) *DB {
	return d
}

// Not filter records that don't match current conditions
func (d *DB) Not(query interface{}, args ...interface{}) *DB {
	return d
}

// Limit specify the number of records to retrieve
func (d *DB) Limit(limit interface{}) *DB {
	return d
}

// Save will insert or update value in database.
func (d *DB) Save(value interface{}) *DB {
	return d
}

// Create insert the value into database
func (d *DB) Create(item interface{}) *DB {
	return d
}

// BeginTransaction (obviously) begins a Transaction
func (d *DB) BeginTransaction() *DB {
	txFunc := func(d *DB, ctx context.Context) *DB {
		dbClone := d.clone()
		// transaction logic
		return dbClone
	}
	return txFunc(d, context.Background())
}

// Commit a Transaction
func (d *DB) Commit() *DB {
	return d
}

// Rollback a Transaction
func (d *DB) Rollback() *DB {
	return d
}

// Transaction starts a database Transaction block
func (d *DB) Transaction() error {
	return nil
}

// helper/private methods

// clone current db connection
func (d *DB) clone() *DB {
	db := &DB{
		db: d.db,
	}
	if db.search == nil {
		db.search = &search{}
	} else {
		db.search = db.search.clone()
	}
	db.search.db = db
	return db
}
