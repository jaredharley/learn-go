package main

import (
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
    _ "reflect"
    "strconv"
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
            userInfo = GetUserInfo(conv)
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
            //var userTasks Task
            userInfo := GetUserInfo(conv)
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

func main() {
    fmt.Println("╔═════════════════╗")
    fmt.Println("║ Welcome to TODO ║")
    fmt.Println("╚═════════════════╝")

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