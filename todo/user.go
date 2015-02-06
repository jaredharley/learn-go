package main

import (
    "database/sql"
    "errors"
    _ "github.com/mattn/go-sqlite3"
)

// User struct
type User struct {
    Id          int
    Firstname   string
    Lastname    string
    Dob         string
    Email       string
}

// Returns the user data from the database for the provided user id.
func GetUserInfo(id int, db *sql.DB) (User, error) {
    var myUser User
    
    // Check to see if the database pointer has a nil reference.
    if db == nil {
        return myUser, errors.New("Unable to connect to the database.")
    }

    stmt, err := db.Prepare("SELECT * FROM users WHERE users.id = ?")
    if err != nil {
        return myUser, err
    }
    defer stmt.Close()

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

    return myUser, err
}

