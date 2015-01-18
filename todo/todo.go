package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "reflect"
    "strconv"
)

func getUserInfo(id int) string {
    fmt.Println("Opening connection to database...")
    fmt.Printf("id is %d, type of %s\n", id, reflect.TypeOf(id).Kind())
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
    var firstname string
    var userid int
    var lastname string
    var dob string
    var email string
    err = stmt.QueryRow(id).Scan(&userid, &firstname, &lastname, &dob, &email)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Printf("User %d is %s\n", id, firstname)

    return "hi"
}

func main() {
    fmt.Println("Welcome to TODO")
    fmt.Println("===============")
    fmt.Printf("Enter a user id: ")
    
    // Read the input
    var text string
    _, err := fmt.Scanf("%s", &text)
    if err != nil {
        log.Fatalln(err)
    }
    
    // Convert the text string into an int
    conv, err := strconv.Atoi(text)
    
    if err != nil {
        log.Fatalln("Uh oh! Couldn't convert the user id.\n%s", err)
        return
    }
    
    // Get the user info
    getUserInfo(conv)
    
}