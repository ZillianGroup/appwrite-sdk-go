package main

import (
	"log"
	"os"
	"time"

	"github.com/zilliangroup/appwrite-sdk-go/appwrite"
)

const (
	EmptyParentDocument     = ""
	EmptyParentProperty     = ""
	EmptyParentPropertyType = ""
)

func main() {
	var EmptyArray = []interface{}{}

	client := appwrite.NewClient()
	client.SetEndpoint(os.Getenv("YOUR_ENDPOINT"))
	client.SetProject(os.Getenv("YOUR_PROJECT_ID"))
	client.SetKey(os.Getenv("YOUR_KEY"))
	client.SetTimeout(10 * time.Second)

	db := appwrite.NewDatabase(client)
	data := map[string]string{
		"hello": "world",
	}
	doc, err := db.CreateDocument(
		os.Getenv("COLLECTION_ID"),
		EmptyParentDocument,
		data,
		EmptyArray,
		EmptyArray,
	)
	if err != nil {
		log.Printf("Error creating document: %v", err)
	}
	log.Printf("Created document: %v", doc)
}
