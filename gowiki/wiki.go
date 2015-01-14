package main

// Imports required modules.
import (
	"fmt"
	"io/ioutil"
	"net/http"
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

// View handler
// This handler extracts the page title from the URL (everything after /view/),
// loads the page from disk, and prints the formatted data to the 
// http.ResponseWriter.
// Currently ignores errors from loadPage.
func viewHandler(w http.ResponseWriter, r *http.Request) {
     title := r.URL.Path[len("/view/"):]
     p, _ := loadPage(title)
     fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// Main function
// Main entry point for the application.
func main() {
     http.HandleFunc("/view/", viewHandler)
     http.ListenAndServe(":8000", nil)
}