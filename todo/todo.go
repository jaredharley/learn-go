package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
    _ "reflect"
    "strconv"
)

// User struct
type User struct {
    id          int
    firstname   string
    lastname    string
    dob         string
    email       string
}

// Task struct
type Task struct {
    id          int
    userId      int
    title       string
    description string
    due_date    string
    importance  int
}

func getUserInfo(id int) User {
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
        myUser.id = id
        myUser.firstname = fn
        myUser.lastname = ln
        myUser.dob = dob
        myUser.email = email
    }

    return myUser
}

func lookupUser() {
Lookup:    
    for {
        fmt.Printf("[User lookup] Enter a user id, (b)ack, or (q)uit: ")
        // Read the input
        var text string
        _, err := fmt.Scanf("%s", &text)
        if err != nil {
            log.Fatalln(err)
        }
        
        switch text {
            case "q":
                // Exit the program
                os.Exit(0);
            case "b":
                break Lookup
        }
        
        // Convert the text string into an int
        conv, err := strconv.Atoi(text)
        
        if err == nil {
            // Get the user info as a User struct
            var userInfo User
            userInfo = getUserInfo(conv)
                // Check and see if the returned struct is equal to an empty struct -
            // this tells us if the returned object was null or not.
            if userInfo == (User{}) {
                fmt.Println("That user does not exist.\n")
            } else {
                fmt.Printf("User %d is %s %s (%s)\n\n", userInfo.id, userInfo.firstname, userInfo.lastname, userInfo.email)
            }
        } else {
            fmt.Printf("%s is not a valid user id\n", text)
        }
    }
}

func lookupTask() {
Lookup:
    for {
        fmt.Printf("[Task lookup] Enter a user id, (b)ack, or (q)uit: ")
        // Read the input
        var text string
        _, err := fmt.Scanf("%s", &text)
        if err != nil {
            log.Fatalln(err)
        }
        
        switch text {
            case "q":
                // Exit the program
                os.Exit(0);
            case "b":
                break Lookup
        }
        
        // Convert the text string into an int
        conv, err := strconv.Atoi(text)
        
        if err == nil {
            // Get the user info as a User struct
            var userTasks Task
            userInfo = getUserInfo(conv)
                // Check and see if the returned struct is equal to an empty struct -
            // this tells us if the returned object was null or not.
            if userInfo == (User{}) {
                fmt.Println("That user does not exist.\n")
            } else {
                fmt.Printf("User %d is %s %s (%s)\n\n", userInfo.id, userInfo.firstname, userInfo.lastname, userInfo.email)
            }
        } else {
            fmt.Printf("%s is not a valid user id\n", text)
        }
    }
}

func main() {
    fmt.Println("Welcome to TODO")
    fmt.Println("===============")
    for {
        fmt.Printf("[Main menu] Select an option (u)sers, (t)asks, (q)uit: ")
        // Read the input
        var text string
        _, err := fmt.Scanf("%s", &text)
        if err != nil {
            log.Fatalln(err)
        }
        
        switch text {
            case "q":
                // Exit the program
                os.Exit(0);
            case "u":
                lookupUser()
            case "t":
                lookupTask()
        }

        
    }
}