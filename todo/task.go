package main

import (
    "database/sql"
    "errors"
	"fmt"
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

// Returns a string array with a list of all tasks assigned to the provided
// user id.
func GetListofTasks(id int, db *sql.DB) ([]string, error) {
    var tasklist []string
    
    if db == nil {
        return tasklist, errors.New("Unable to connect to the database")
    }
    
    query := `SELECT task_id,
                     user_id,
                     title,
                     due_date,
                     importance
                FROM tasks
               WHERE tasks.user_id = ?`

    rows, err := db.Query(query, id)
    if err != nil {
        return tasklist, err
    }
    defer rows.Close()
    for rows.Next() {
    	var task_id int
    	var user_id int
    	var title string
    	var due_date sql.NullString
    	var importance int
    	if err := rows.Scan(&task_id, &user_id, &title, &due_date, &importance); err != nil {
    		return tasklist, err
    	}

    	date := ""
    	if due_date.Valid {
    	    date = due_date.String
    	}
    	
    	item := fmt.Sprintf("%d: %s (%s){%d}", task_id, title, date, importance)
    	tasklist = append(tasklist, item)
    }

    return tasklist, err

}

