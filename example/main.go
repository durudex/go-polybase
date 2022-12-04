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
	coll := polybase.New(polybase.Config{
		URL: polybase.TestnetURL,
	}).Collection("V-space/City")

	getSingle(coll)

	getList(coll)

	create(coll)
}

func getSingle(coll polybase.Collection) {
	var single polybase.SingleResponse[City]

	coll.Record("Hello Durudex").Get(context.Background(), &single)

	fmt.Println("Single, Block Hash:", single.Block.Hash)
	fmt.Println("Single, Data:", single.Data.ID, single.Data.Name)
}

func getList(coll polybase.Collection) {
	var result polybase.Response[City]

	coll.Get(context.Background(), &result)

	for _, city := range result.Data {
		fmt.Println("List, Block Hash:", city.Block.Hash)
		fmt.Println("List, Data:", city.Data.ID, city.Data.Name)
	}

	fmt.Println("List, After Cursor:", result.Cursor.After)
	fmt.Println("List, Before Cursor:", result.Cursor.Before)
}

func create(coll polybase.Collection) {
	coll.Create(context.Background(), []any{"Hello Durudex", "Hello Durudex"})
}
