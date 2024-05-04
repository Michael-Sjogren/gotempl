package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(dbname string) (*sql.DB, error) {
	return sql.Open("sqlite3", dbname)
}

func InitializeDB(con *sql.DB) error {
	res, err := con.Exec(`
	CREATE TABLE IF NOT EXISTS users 
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password_hash BLOB NOT NULL,
		created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		access INTEGER NOT NULL DEFAULT 0
	) STRICT;`)

	if err != nil {
		return err
	}

	n, err := res.RowsAffected()

	if err == nil {
		log.Printf("Rows affected: %d\n", n)
	}

	return nil
}
