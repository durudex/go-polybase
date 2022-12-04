package main

import (
	"context"
	"fmt"

	"github.com/v1def/go-polybase"
)

type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	coll := polybase.New().Collection("V-space/City")

	var result polybase.SingleResponse[City]

	coll.Record("Vizag").Get(context.Background(), &result)

	fmt.Println(result.Block.Hash)
	fmt.Println(result.Data.Name)
}
