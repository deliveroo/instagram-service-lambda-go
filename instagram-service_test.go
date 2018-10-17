package main

import (
    "testing"
    "github.com/aws/aws-lambda-go/events"
)

func TestFound(t *testing.T) {
    //When the restaurant exists
    expected := events.APIGatewayProxyResponse{Body: "{\"handle\":\"nandosuk\"}", StatusCode: 200}  
    var query = map[string]string{ "restaurantid": "14950", }
    req := events.APIGatewayProxyRequest{QueryStringParameters: query}
    var actual, err = HandleRequest(req)
    if actual.Body != expected.Body {
        t.Errorf("Error actual: '%v', and Expected: '%v'.", actual, expected)
    }
    if err != nil {
        t.Errorf("Error: %v", err)
    }
}

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