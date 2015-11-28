package main

// Imports required modules.
import (
    "errors"
    "flag"
    "fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

// Page struct
// Builds a struct defining page objects.
type Page struct {
     Title string
     Body  []byte
     FileList []string
}


//// Global vars
var (
    addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)
var td = "templates/"
var templates = template.Must(template.ParseFiles(td + "edit.html", 
                                                  td + "view.html",
                                                  td + "main.html",))
var validPath = regexp.MustCompile("^/console/(edit|save|view)/(.+)$")
var fileDir = "/Users/jared/test.jaredharley.com/_posts/"

// Save function
// Saves the page to disk, using the page title as the filename, and writes
// the file using read/write permissions for the current user only (0600)
func (p *Page) save() error {
     filename := fileDir + p.Title
     fmt.Println("Saving " + filename)
     return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load function
// Loads the page from disk, using the title to find the file on disk. The
// loadPage function returns a pointer to the Page literal, or an error if
// an error occurred.
func loadPage(title string) (*Page, error) {
     filename := fileDir + title
     fmt.Println("Reading " + filename)
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
    fmt.Println("in viewHandler")
     p, err := loadPage(title)
     if err != nil {
         http.Redirect(w, r, "/console/edit/" + title, http.StatusFound)
         return
     }
     renderTemplate(w, "view", p)
}

// Edit handler
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    fmt.Println("In editHandler()")
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

// Save handler
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    fmt.Println("In saveHandler()")
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/console/view/" + title, http.StatusFound)
}

// Main page handler
func mainPageHandler(w http.ResponseWriter, r *http.Request, title string) {
     p := &Page{FileList: buildFileList(fileDir)}
     renderTemplate(w, "main", p)
}

func buildFileList(dir string) []string {
    var fileList []string
    directory, err := os.Open(dir)
    if err != nil {
        return nil
    }
    defer directory.Close()
    
    fileInfo, err := directory.Readdir(-1)
    if err != nil {
        return nil
    }
    for _, fi := range fileInfo {
        fileList = append(fileList, fi.Name())
    }
    return fileList
}

// Template renderer
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    fmt.Println(p)
    err := templates.ExecuteTemplate(w, tmpl + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// Get title
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    fmt.Println("in getTitle()")
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Invalid Page Title")
    }
    
    // The title is in the second subexpression
    return m[2], nil
}

// buildBlogHandler
func buildBlogHandler(w http.ResponseWriter, r *http.Request, title string) {
    fmt.Println("============================")
    fmt.Println("Begin blog build")
    jekyll_out, err := exec.Command("/home/jared/test-buildsite.sh").Output()
    fmt.Println(string(jekyll_out))
    copy_out, err := exec.Command("/usr/bin/sudo","/home/jared/test-copysite.sh").Output()
    fmt.Println(string(copy_out))
    fmt.Println("End blog build")
    fmt.Println("============================")
    if err != nil {
        fmt.Println("ERROR")
        fmt.Println(err)
    }
    http.Redirect(w, r, "/console/", http.StatusFound)
}

// Build handler function
func buildHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("URL requested: " + r.URL.Path)
        if r.URL.Path == "/console/" {
            fn(w, r, "/console/")
        } else if r.URL.Path == "/console/build/" {
            fn(w, r, "/console/build/")
        } else {
            m := validPath.FindStringSubmatch(r.URL.Path)
            if m == nil {
                fmt.Println("Not found.")
                http.NotFound(w, r)
                return
            }
                fn(w, r, m[2])
        }
    }
}

// Main function
// Main entry point for the application.
func main() {
    
    flag.Parse()
    
    http.HandleFunc("/console/", buildHandler(mainPageHandler))
    http.HandleFunc("/console/view/", buildHandler(viewHandler))
    http.HandleFunc("/console/edit/", buildHandler(editHandler))
    http.HandleFunc("/console/save/", buildHandler(saveHandler))
    http.HandleFunc("/console/build/", buildHandler(buildBlogHandler))
    
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
    
    fmt.Println("Starting server")
    http.ListenAndServe(":8000", nil)
}