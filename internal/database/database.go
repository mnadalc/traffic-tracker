package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	conn *sql.DB
}

func Connect() (*DB, error) {
	conn, err := sql.Open("sqlite", "traffic.db")

	if err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}

// Execute routes and traffic data table creation
func (db *DB) Init() error {
	_, err := db.conn.Exec(createRoutesTable)
	if err != nil {
		return err
	}
	
	_, err = db.conn.Exec(createTrafficDataTable)
	if err != nil {
		return err
	}
	
	return nil
}