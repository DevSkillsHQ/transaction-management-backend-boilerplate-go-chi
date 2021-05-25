package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type State struct {
	DB        *sql.DB
	WriteLock sync.Mutex
}

func Init(path string, reinit bool) *State {

	if reinit {
		// Reset with every run for testing purposes
		err := os.Remove("db.sqlite")
		if err != nil {
			log.Fatal("Cannot create db table")
		}
	}

	var dbs State
	dbs.DB, _ = sql.Open("sqlite3", "db.sqlite")

	if reinit {
		dbs.prepareDatabase()
	}

	return &dbs
}

func (db *State) Close() {
	db.DB.Close()
}

func (db *State) Lock() {
	db.WriteLock.Lock()
}

func (db *State) Unlock() {
	db.WriteLock.Unlock()
}

func (db *State) prepareDatabase() {

	// TODO Prepare your DB
	stmt, _ := db.DB.Prepare(`SELECT 1;`)

	_, err := stmt.Exec()
	if err != nil {
		log.Fatal("failed to run the query")
	}

	fmt.Println("the DB is ready")
}
