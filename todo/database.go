package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
)

func OpenDatabaseConnection() *sql.DB {
    db, err := sql.Open("sqlite3", "./todo-app.db")
    if err != nil {
        fmt.Println("Unable to open database")
        // log.Fatal() prints the error and exits the program (equiv to calling
        // os.Exit(1).
        log.Fatalln(err)
    }
    return db
}

func CloseDatabaseConnection(db *sql.DB) {
    err := db.Close()
    if err != nil {
        fmt.Println("Unable to close database")
        // log.Fatal() prints the error and exits the program (equiv to calling
        // os.Exit(1).
        log.Fatalln(err)
    }
}