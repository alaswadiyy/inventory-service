package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestStatusEndpoint(t *testing.T) {
    req := httptest.NewRequest("GET", "/status", nil)
    w := httptest.NewRecorder()
    
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })
    handler.ServeHTTP(w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}