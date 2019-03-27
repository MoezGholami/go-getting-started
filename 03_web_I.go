package main

import (
    "fmt"
    "net/http"
)

const htmlPrefix string = "<html><body>"
const htmlSuffix string = "</body></html>"

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, htmlPrefix+"hello world! I am index. Want to learn <a href=\"/about\">about me</a>?"+htmlSuffix)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, htmlPrefix+"I'm cool. <a href=\"/\">Here is your way to home</a>"+htmlSuffix)
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/about", aboutHandler)
    fmt.Println("going to serve at http://localhost:4000 ")
    http.ListenAndServe(":4000", nil)
    fmt.Println("yes! serve is a blocking operation :(")
}
