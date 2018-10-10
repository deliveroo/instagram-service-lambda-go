package main

import (
    "testing"
    "net/http"
    "fmt"
)

func TestHandler(t *testing.T) {
    req, _ := http.NewRequest("GET", "http://example.com", nil)
    var string, error = HandleRequest(*req)
    fmt.Println(string, error)
    //t.Errorf(result)
}