package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type AboutMePageContent struct {
    Title string
    Adjective string
}

const htmlPrefix string = "<html><body>"
const htmlSuffix string = "</body></html>"

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    content := AboutMePageContent{Title: "About Me", Adjective: "cool"}
    t, _ := template.ParseFiles("09_web_template.html")
    t.Execute(w, content)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, htmlPrefix+"hello world! I am index. Want to learn <a href=\"/about\">about me</a>?"+htmlSuffix)
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/about", aboutHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web_static_files"))))

    fmt.Println("going to serve at http://localhost:4000 ")
    http.ListenAndServe(":4000", nil)
    fmt.Println("yes! serve is a blocking operation :(")
}
