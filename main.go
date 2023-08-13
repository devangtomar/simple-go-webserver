package main

import (
    "fmt"
    "log"
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        // Handle form submission here
        // ...
        fmt.Fprintf(w, "Form Submitted Successfully!")
        return
    }
    http.ServeFile(w, r, "./static/form.html")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/hello.html")
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)

    fmt.Println("Server is running at :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
