package main

import (
	"context"
	"fmt"

	"github.com/durudex/go-polybase"
)

type Collection struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

func main() {
	client := polybase.New(&polybase.Config{
		URL: polybase.TestnetURL,
	})
	coll := polybase.NewCollection[Collection](client, "Collection")

	response := coll.Record("Collection").Get(context.Background())

	fmt.Println("Block Hash:", response.Block.Hash)
	fmt.Println("ID:", response.Data.ID)
	fmt.Println("Code:", response.Data.Code)
}
