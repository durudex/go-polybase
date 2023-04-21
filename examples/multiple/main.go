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

	response := coll.Get(context.Background())

	fmt.Println("After Cursor", response.Cursor.After)
	fmt.Println("Before Cursor", response.Cursor.Before)

	for _, record := range response.Data {
		fmt.Println("Block Hash:", record.Block.Hash)
		fmt.Println("ID:", record.Data.ID)
		fmt.Println("Code:", record.Data.Code)
	}
}
