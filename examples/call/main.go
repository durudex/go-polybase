package main

import (
	"context"
	"fmt"

	"github.com/durudex/go-polybase"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Completed int    `json:"completed"`
}

func main() {
	client := polybase.New(&polybase.Config{
		URL: polybase.TestnetURL,
	})
	coll := polybase.NewCollection[Todo](client, "polybase/todo")

	args := []any{"1", "go-polybase", "example", 0}

	response := coll.Record("1").
		Call(context.Background(), "update", args)

	fmt.Println("ID:", response.Data.ID)
	fmt.Println("Title:", response.Data.Title)
	fmt.Println("Content:", response.Data.Content)
	fmt.Println("Completed:", response.Data.Completed == 1)
}
