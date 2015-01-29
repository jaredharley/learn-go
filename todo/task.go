package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
)

// Task struct
type Task struct {
    Id          int
    UserId      int
    Title       string
    Description string
    Due_date    string
    Importance  int
}

func GetListofTasks(id int) []string {
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

    rows, err := db.Query("SELECT * FROM tasks WHERE tasks.user_id = ?", id)
    if err != nil {
        log.Fatalln(err)
    }
    defer rows.Close()
    var tasklist []string
    for rows.Next() {
    	var task_id int
    	var user_id int
    	var title string
    	var description string
    	var due_date string
    	var importance int
    	if err := rows.Scan(&task_id, &user_id, &title, &description, &due_date, &importance); err != nil {
    		log.Fatalln(err)
    	}
    	fmt.Printf("%d: %s (%s){%d}\n", task_id, title, due_date, importance)
    }

    return tasklist

}