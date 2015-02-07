package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
    _ "reflect"
    "strconv"
    "strings"
)

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
            userInfo, err = GetUserInfo(conv, db)
            
            if err != nil {
                fmt.Println(err)
            }
            
            // Check and see if the returned struct is equal to an empty struct -
            // this tells us if the returned object was null or not.
            if userInfo == (User{}) {
                fmt.Println("That user does not exist.\n")
            } else {
                fmt.Printf("User %d is %s %s (%s)\n\n", userInfo.Id, userInfo.Firstname, userInfo.Lastname, userInfo.Email)
            }
        } else {
            fmt.Printf("%s is not a valid user id\n", text)
        }
    }
}

func lookupTask() {
Lookup:
    for {
        fmt.Printf("\n[Task lookup] Enter a user id, (b)ack, or (q)uit: ")
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
        
        if err != nil {
            fmt.Printf("%s is not a valid user id\n", text)
            continue
        }

        // Get the user info as a User struct
        userInfo, _ := GetUserInfo(conv, db)

        // Check and see if the returned struct is equal to an empty struct -
        // this tells us if the returned object was null or not.
        if userInfo == (User{}) {
            fmt.Println("That user does not exist.\n")
            continue
        }
        titleString := "Tasks for " + userInfo.Firstname + " (" + userInfo.Email + "):"
        fmt.Println(titleString)
        fmt.Println(strings.Repeat("-",len(titleString)))

        taskList, err := GetListofTasks(conv, db)
        if err != nil {
            fmt.Println("Unable to retrieve list of tasks: ")
            fmt.Println(err)
        }
        for i := 0; i < len(taskList); i++ {
            fmt.Println(taskList[i])
        }

    }
}

var db *sql.DB

func main() {
    fmt.Println("╔═════════════════╗")
    fmt.Println("║ Welcome to TODO ║")
    fmt.Println("╚═════════════════╝")

    db = OpenDatabaseConnection()

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
                CloseDatabaseConnection(db)
                os.Exit(0);
            case "u":
                lookupUser()
            case "t":
                lookupTask()
        }

        
    }
}