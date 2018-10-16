package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"
    "net/http"
)

type InvalidRequestError struct {
    Type string
    HttpStatus int
    What string
}

func (e InvalidRequestError) Error() string {
    return fmt.Sprintf("%v: %v: %v", e.Type, e.HttpStatus, e.What)
}

func idNotFoundError(id string) error {
    return InvalidRequestError {
        "Not Found",
        404,
        fmt.Sprintf("Instagram details are not found for %v", id),
    }
}

func invalidRequestError(request http.Request) error {
    return InvalidRequestError {
        "Bad Request",
        400,
        fmt.Sprintf("Request url invalid: %v", request.URL.String()),
    }
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

func HandleRequest(request events.APIGatewayProxyRequest) (string, error) {

    fmt.Println(request)
    restaurantID := request.QueryStringParameters["restaurantid"]

    fmt.Println(request.QueryStringParameter)

    if restaurantID == "" {
        fmt.Println("NO restaurantid QUERY PARAM")
        //return "", invalidRequestError(request)
    }

    instagramHandle := instagramHandlesById[restaurantID]
    if instagramHandle == "" {
        fmt.Println("NO INSTAGRAM HANDLE")
        //return "", idNotFoundError(restaurantID)
    }

    return instagramHandle, nil
}
