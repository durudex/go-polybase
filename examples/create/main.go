package main

import (
	"context"
	"fmt"
	"time"

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

	input := &Todo{
		ID:        time.Now().String(),
		Title:     "go-polybase",
		Content:   "example",
		Completed: 0,
	}

	response := coll.Create(context.Background(), input)

	fmt.Println("ID:", response.Data.ID)
	fmt.Println("Title:", response.Data.Title)
	fmt.Println("Content:", response.Data.Content)
	fmt.Println("Completed:", response.Data.Completed == 1)
}
