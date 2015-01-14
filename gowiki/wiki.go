package main

// Imports required modules.
import (
    "errors"
    "flag"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

// Page struct
// Builds a struct defining page objects.
type Page struct {
     Title string
     Body  []byte
}

//// Global vars
var (
    addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)
var td = "templates/"
var templates = template.Must(template.ParseFiles(td + "edit.html", 
                                                  td + "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// Save function
// Saves the page to disk, using the page title as the filename, and writes
// the file using read/write permissions for the current user only (0600)
func (p *Page) save() error {
     filename := "pages/" + p.Title + ".txt"
     return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load function
// Loads the page from disk, using the title to find the file on disk. The
// loadPage function returns a pointer to the Page literal, or an error if
// an error occurred.
func loadPage(title string) (*Page, error) {
     filename := "pages/" + title + ".txt"
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
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
     p, err := loadPage(title)
     if err != nil {
         http.Redirect(w, r, "/edit/" + title, http.StatusFound)
         return
     }
     renderTemplate(w, "view", p)
}

// Edit handler
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

// Save handler
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

// Template renderer
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// Get title
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Invalid Page Title")
    }
    
    // The title is in the second subexpression
    return m[2], nil
}


// Build handler function
func buildHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}

// Main function
// Main entry point for the application.
func main() {
    
    flag.Parse()
    
    http.HandleFunc("/view/", buildHandler(viewHandler))
    http.HandleFunc("/edit/", buildHandler(editHandler))
    http.HandleFunc("/save/", buildHandler(saveHandler))
    
    if *addr {
        l, err := net.Listen("tcp", "127.0.0.1:0")
        if err != nil {
            log.Fatal(err)
        }
        err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
        if err != nil {
            log.Fatal(err)
        }
        s := &http.Server{}
        s.Serve(l)
        return
    }
    
    http.ListenAndServe(":8000", nil)
}