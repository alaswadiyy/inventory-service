package main

import (
    "encoding/json"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
        items := []string{"Laptop", "Desktop", "Tablet"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(items)
    })

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}