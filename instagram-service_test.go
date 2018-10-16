package main

// import (
//     "testing"
//     "net/http"
// )

// func TestFound(t *testing.T) {
//     //When the restaurant exists
//     Expected := "nandosuk"
//     req, _ := http.NewRequest("GET", "http://example.com/?restaurantID=14950", nil)
//     var actual, err = HandleRequest(*req)
//     if err != nil {
//         t.Errorf("Error: %v", err)
//     } else if actual != Expected {
//         t.Errorf("Error actual: '%v', and Expected: '%v'.", actual, Expected)
//     }
// }

// func TestNotFound(t *testing.T) {
//     //When the restaurant ID doesn't exist
//     Expected := idNotFoundError("99999999999")
//     req, _ := http.NewRequest("GET", "http://example.com/?restaurantID=99999999999", nil)
//     var _, err = HandleRequest(*req)
//     if err != Expected {
//         t.Errorf("Actual error: %v, Expected Error: %v", err.Error(), Expected.Error())
//     }
// } 

// func TestInvalidRequest(t *testing.T) {
//     //When the restaurant ID doesn't exist
//     req, _ := http.NewRequest("GET", "http://example.com/?resantID=14950", nil)
//     Expected := invalidRequestError(*req)
//     var _, err = HandleRequest(*req)
//     if err != Expected {
//         t.Errorf("Actual error: %v, Expected Error: %v", err.Error(), Expected.Error())
//     }
// }