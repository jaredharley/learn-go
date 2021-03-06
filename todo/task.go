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
               WHERE tasks.user_id = ?;`

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

// This function creates a new task for the specified user, using the provided
// Task struct, and returns an error if it fails, and the id of the new task
// if successful.
func CreateNewTask(userid int, newTask Task, db *sql.DB) (int64, error) {
    var newTaskId int64 = -1
    // Validate user id, error if not valid.
    validUser, err := IsUserValid(userid, db)
    if !validUser {
        errString := fmt.Errorf("Invalid user id. %v", err)
        return newTaskId, errString
    }
    
    // Validate manditory task values, error if not valid.
    // Required: user_id, title
    if newTask.Title == "" {
        return newTaskId, fmt.Errorf("Title is a required field.")
    }
    
    query := `INSERT INTO tasks (
                user_id,
                title,
                description,
                due_date,
                importance
              ) values (
                ?,
                ?,
                ?,
                ?,
                ?
              );`
    
    addTask, err := db.Prepare(query)
    if err != nil {
        return newTaskId, err
    }
    
    // Start a new transaction
    tx, err := db.Begin()
    if err != nil {
        tx.Rollback()
        return newTaskId, err
    }
    
    result, err := tx.Stmt(addTask).Exec(userid, newTask.Title, newTask.Description, newTask.Due_date, newTask.Importance)
    if err != nil {
        tx.Rollback()
        return newTaskId, err
    } else {
        tx.Commit()
        newTaskId, _ = result.LastInsertId()
    }
    
    return newTaskId, nil
}