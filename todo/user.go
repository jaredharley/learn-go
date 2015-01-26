package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
)

// User struct
type User struct {
    Id          int
    Firstname   string
    Lastname    string
    Dob         string
    Email       string
}

func GetUserInfo(id int) User {
    //fmt.Println("Opening connection to database...")
    //fmt.Printf("id is %d, type of %s\n", id, reflect.TypeOf(id).Kind())
    db, err := sql.Open("sqlite3", "./todo-app.db")
    if err != nil {
        fmt.Println("Unable to open database")
        // log.Fatal() prints the error and exits the program (equiv to calling
        // os.Exit(1).
        log.Fatalln(err)
    }
    
    // Deferring the database close - ensures the db.Close() call will
    // be called as the main() function finishes.
    defer db.Close()
    
    stmt, err := db.Prepare("SELECT * FROM users WHERE users.id = ?")
    if err != nil {
        log.Fatalln(err)
    }
    defer stmt.Close()
    var myUser User

    var uid     int
    var fn      string
    var ln      string
    var dob     string
    var email   string
    err = stmt.QueryRow(id).Scan(&uid, &fn, &ln, &dob, &email)
    if err == nil {
        myUser.Id = id
        myUser.Firstname = fn
        myUser.Lastname = ln
        myUser.Dob = dob
        myUser.Email = email
    }

    return myUser
}