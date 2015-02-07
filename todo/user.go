package main

import (
    "database/sql"
    "errors"
)

// User struct
type User struct {
    Id          int
    Firstname   string
    Lastname    string
    Dob         string
    Email       string
}

const getUserInfoSelect string = `SELECT id,
                                         first_name,
                                         last_name,
                                         dob,
                                         email
                                    FROM users
                                   WHERE users.id = ?`

// Returns the user data from the database for the provided user id.
func GetUserInfo(id int, db *sql.DB) (User, error) {
    var myUser User
    
    // Check to see if the database pointer has a nil reference.
    if db == nil {
        return myUser, errors.New("Unable to connect to the database.")
    }

    stmt, err := db.Prepare(getUserInfoSelect)
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

