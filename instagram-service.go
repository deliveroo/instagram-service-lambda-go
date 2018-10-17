package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"
)

func idNotFoundErrorResponse(id string) events.APIGatewayProxyResponse {
    body:= fmt.Sprintf("{\"error\":\"Instagram details are not found for %v\"}", id)
    return events.APIGatewayProxyResponse{Body: body, StatusCode: 404}
}

func invalidRequestErrorResponse() events.APIGatewayProxyResponse {
    body:= fmt.Sprintf("{\"error\":\"Request url invalid\"}")
    return events.APIGatewayProxyResponse{Body: body, StatusCode: 400}
}

var responseHeaders = map[string]string {
    "cache-control": "public, max-age:86400",
}

var instagramHandlesById map[string]string

func init() {
    var jsonFile, jsonfileError = ioutil.ReadFile("data.json")
        if jsonfileError != nil {
        fmt.Printf("File error: %v\n", jsonfileError)
    }
    json.Unmarshal([]byte(jsonFile), &instagramHandlesById)
}

func main() {
    lambda.Start(HandleRequest)
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    restaurantID := request.QueryStringParameters["restaurantid"]

    if restaurantID == "" {
        return invalidRequestErrorResponse(), nil
    }

    instagramHandle := instagramHandlesById[restaurantID]
    if instagramHandle == "" {
        return idNotFoundErrorResponse(restaurantID), nil
    }

    body := fmt.Sprintf("{\"url\":\"https://www.instagram.com/%s/\"}", instagramHandle)
    return events.APIGatewayProxyResponse{Headers: responseHeaders, Body: body, StatusCode: 200}, nil
}
