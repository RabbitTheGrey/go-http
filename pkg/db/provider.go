package db

import "database/sql"

var db *sql.DB

func SetDB(database *sql.DB) {
    db = database
}

func GetDB() *sql.DB {
    if db == nil {
        panic("DB not initialized")
    }
    return db
}
