package database

import (
	"database/sql"
	"fmt"
)

type Db struct {
	db *sql.DB
}

func (db *Db) NewConnector() error {
	// sqlite3 in memory
	var err error
	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", "test.db")
	db.db, err = sql.Open("sqlite", dsn)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) InitDb() {
	query := `CREATE TABLE IF NOT EXISTS authors (
    			id INTEGER PRIMARY KEY AUTOINCREMENT,
    			name TEXT NOT NULL)`
	db.db.Exec(query)
}

func (db *Db) SaveAuthor(name string) error {
	query := `INSERT INTO authors (name) VALUES (?)`
	result, err := db.db.Exec(query, name)
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("result is nil")
	}
	return nil
}

func (db *Db) Get
