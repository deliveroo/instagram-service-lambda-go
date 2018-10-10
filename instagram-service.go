package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "github.com/aws/aws-lambda-go/lambda"
    "net/http"
)

// type Instagram struct {
//   Name string
// }

// type MyEvent struct {
//         Name string `json:"name"`
// }

var instagramHandlesById map[string]string

func init() {
    var jsonFile, jsonfileError = ioutil.ReadFile("./data.json")
        if jsonfileError != nil {
        fmt.Printf("File error: %v\n", jsonfileError)
    }
    json.Unmarshal([]byte(jsonFile), &instagramHandlesById)
}

func main() {
    lambda.Start(HandleRequest)
}

func HandleRequest(request http.Request) (string, error) {
    for key, value := range instagramHandlesById {
        fmt.Println(key, value)
    }
    return instagramHandlesById["14950"], nil
}

func getInstagramHandleForID(id string) (string, bool) {
    var thing, thing2 = instagramHandlesById[id]
    return thing, thing2
}