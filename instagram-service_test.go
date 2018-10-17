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
    var actual = HandleRequest(req)
    if actual.Body != expected.Body {
        t.Errorf("Error actual: '%v', and Expected: '%v'.", actual, expected)
    }
}

func TestNotFound(t *testing.T) {
    //When the restaurant ID doesn't exist
    expected := idNotFoundErrorResponse("99999999999")
    var query = map[string]string{ "restaurantid": "99999999999", }
    req := events.APIGatewayProxyRequest{QueryStringParameters: query}
    var actual = HandleRequest(req)
    if actual.Body != expected.Body {
        t.Errorf("Error actual: '%v', and Expected: '%v'.", actual, expected)
    }
} 

func TestInvalidRequest(t *testing.T) {
    //When the query param is wrong
    expected := invalidRequestErrorResponse()
    var query = map[string]string{ "resantID": "14950", }
    req := events.APIGatewayProxyRequest{QueryStringParameters: query}
    var actual = HandleRequest(req)
    if actual.Body != expected.Body {
        t.Errorf("Error actual: '%v', and Expected: '%v'.", actual, expected)
    }
}