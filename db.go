package main

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

//LoadDatabase - Load hits db
func LoadDatabase() *sql.DB {
	dbPath := "./database"
	_ = os.Mkdir(dbPath, 0700)

	db, err := sql.Open("sqlite", dbPath+"/hits.db")
	if err != nil {
		log.Fatal(err)
	}

	loadStmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS HITS (id INTEGER PRIMARY KEY, method TEXT, user_agent TEXT, host TEXT, lang TEXT, encoding TEXT)
	`)

	loadStmt.Exec()

	return db
}
