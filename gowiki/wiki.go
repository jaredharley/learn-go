package main

// Imports required modules.
import (
	"fmt"
	"io/ioutil"
)

// Page struct
// Builds a struct defining page objects.
type Page struct {
     Title string
     Body  []byte
}

// Save function
// Saves the page to disk, using the page title as the filename, and writes
// the file using read/write permissions for the current user only (0600)
func (p *Page) save() error {
     filename := p.Title + ".txt"
     return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load function
// Loads the page from disk, using the title to find the file on disk. The
// loadPage function returns a pointer to the Page literal, or an error if
// an error occurred.
func loadPage(title string) (*Page, error) {
     filename := title + ".txt"
     body, err := ioutil.ReadFile(filename)
     if err != nil {
          return nil, err
     }
     return &Page{Title: title, Body: body}, nil
}

// Main function
// Main entry point for the application.
func main() {
     p1 := &Page{Title: "TestPage", Body: []byte("This is my test page")}
     p1.save()
     p2, _ := loadPage("TestPage")
     fmt.Println(string(p2.Body))
}