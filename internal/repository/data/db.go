package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

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

func Create(db *sql.DB) error {
	log.Print("Start table creation")

	usr_query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		type TEXT NOT NULL CHECK(type IN ("Requester", "Responder")),
		is_active BOOLEAN NOT NULL,
		created_at TEXT NOT NULL
	);`

	rqst_tp_query := `CREATE TABLE IF NOT EXISTS request_types(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	);`

	rqst_query := `
	PRAGMA foreign_keys = ON;
	CREATE TABLE IF NOT EXISTS requests (
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		created_at INTEGER DEFAULT CURRENT_TIMESTAMP,
		closed_at INTEGER,

		registrator_id INTEGER,
		responder_id INTEGER,
		request_type_id INTEGER,

		FOREIGN KEY(registrator_id) REFERENCES users(id)
		FOREIGN KEY(responder_id) REFERENCES users(id)
		FOREIGN KEY(request_type_id) REFERENCES request_types(id)
	);`

	_, err := db.Exec(usr_query)
	_, err = db.Exec(rqst_query)
	_, err = db.Exec(rqst_tp_query)

	if err != nil {
		log.Print("Unable to create tables: ", err)
		return err
	}
	return nil
}
