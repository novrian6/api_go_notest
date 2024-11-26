// main.go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Updated World!")  // Update the response
}

func main() {
    port := ":8070"
    fmt.Printf("Starting server on port%s\n", port)
    http.HandleFunc("/", handler)
    http.ListenAndServe(port, nil)
    fmt.Printf("Application stopped\n")
}