package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//LoadDatabase - Load hits db
func LoadDatabase(rootDir string) *sql.DB {
	dbPath := rootDir + "/database"
	_ = os.Mkdir(dbPath, 0700)

	db, err := sql.Open("sqlite3", dbPath+"/hits.db")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`
		CREATE TABLE IF NOT EXISTS HITS (id INTEGER PRIMARY KEY, method TEXT, user_agent TEXT, host TEXT, lang TEXT, encoding TEXT)
	`)

	return db
}
