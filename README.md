# Appwrite Go SDK

![License](https://img.shields.io/github/license/appwrite/sdk-for-go.svg?style=flat-square)
![Version](https://img.shields.io/badge/api%20version-1.2.1-blue.svg?style=flat-square)
[![Build Status](https://img.shields.io/travis/com/appwrite/sdk-generator?style=flat-square)](https://travis-ci.com/appwrite/sdk-generator)

**WORK IN PROGRESS - NOT READY FOR USAGE**

This repository is maintained solely to support [illa-builder](https://github.com/illacloud/illa-builder). It should not be used in production environments.

## Installation

To install using `go get`:

```bash
go get github.com/illacloud/appwrite-sdk-go
```

## Testing the SDK

* clone this repo.
* create a project and within this project a collection.
* configure the documents in the collection to have a _key_ = __hello__.
* Then inject these environment variables:

  ```bash
  export YOUR_ENDPOINT=https://appwrite.io/v1  
  export YOUR_PROJECT_ID=6…8  
  export YOUR_KEY="7055781…cd95"  
  export COLLECTION_ID=616a095b20180  
  ```

Create `main.go` file with:

```go
package main

import (
	"log"
	"os"
	"time"

	"github.com/illacloud/appwrite-sdk-go/appwrite"
)

func main() {
	client := appwrite.NewClient(10 * time.Second)
	client.SetEndpoint(os.Getenv("YOUR_ENDPOINT"))
	client.SetProject(os.Getenv("YOUR_PROJECT_ID"))
	client.SetKey(os.Getenv("YOUR_KEY"))

	db := appwrite.NewDatabase(client)
	data := map[string]string{
		"hello": "world",
	}
	var EmptyArray = []interface{}{}
	doc, err := db.CreateDocument(
		os.Getenv("COLLECTION_ID"),
		data,
		EmptyArray,
		EmptyArray,
		"",
		"",
		"",
	)
	if err != nil {
		log.Printf("Error creating document: %v", err)
	}
	log.Printf("Created document: %v", doc)
}
```

* After that, run the following

  > % go run main.go  
  > 2021/10/16 03:41:17 Created document: map[$collection:616a095b20180 $id:616a2dbd4df16 $permissions:map[read:[] write:[]] hello:world]  


## Contribution

This library is auto-generated by Appwrite custom [SDK Generator](https://github.com/appwrite/sdk-generator). This repository is maintained solely to support [illa-builder](https://github.com/illacloud/illa-builder). External contributions are not accepted at this time. To learn more about how you can help us improve this SDK, please check the Appwrite SDK Generator [contribution guide](https://github.com/appwrite/sdk-generator/blob/master/CONTRIBUTING.md).

## License

Please see the [BSD-3-Clause license](https://raw.githubusercontent.com/appwrite/appwrite/master/LICENSE) file for more information.