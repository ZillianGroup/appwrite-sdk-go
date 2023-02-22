package main

import (
    "fmt"
    "time"
    "github.com/illacloud/appwrite-sdk-go/appwrite"
)

func main() {
    client := appwrite.NewClient()

    client.SetEndpoint("https://[HOSTNAME_OR_IP]/v1") // Your API Endpoint
    client.SetProject("") // Your project ID
    client.SetKey("") // Your secret API key

    var service := appwrite.Databases{
        client: &client
    }

    var response, error := service.CreateDocument("[DATABASE_ID]", "[COLLECTION_ID]", "[DOCUMENT_ID]", nil, ["read("any")"])

    if error != nil {
        panic(error)
    }

    fmt.Println(response)
}
