package data

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	DbPath       string
	MaxOpenConns int
}

func (conn *SqliteDB) Initialize(scriptPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", conn.DbPath+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conn.MaxOpenConns)

	bytes, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Println("Unable to read init.sql file ", err)
		return db, errors.New("Unable to read init db file " + err.Error())
	}

	_, err = db.Exec(string(bytes))
	if err != nil {
		return db, errors.New("Unable to execute entities initialization script init.sql " + err.Error())
	}

	return db, nil
}
