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

	var single polybase.SingleResponse[City]

	coll.Record("Vizag").Get(context.Background(), &single)

	fmt.Println(single.Block.Hash)
	fmt.Println(single.Data.Name)

	var result polybase.Response[City]

	coll.Get(context.Background(), &result)

	fmt.Println(result.Cursor.After)
	fmt.Println(result.Data[0].Data.Name)
}
