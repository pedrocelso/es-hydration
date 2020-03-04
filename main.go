package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println("======== BEGIN =======")
	spew.Dump(res)
	fmt.Println("======= END ========")

	fmt.Println("vim-go")
}
