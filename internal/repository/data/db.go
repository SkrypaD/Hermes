package data

import "database/sql"

type SqliteDB struct {
	DbPath       string
	MaxOpenConns int
}

func (conn *SqliteDB) Initialize() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", conn.DbPath)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conn.MaxOpenConns)

	return db, nil
}
