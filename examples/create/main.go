/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

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

	response := coll.Create(context.Background(), polybase.ParseInput(input))

	fmt.Println("ID:", response.Data.ID)
	fmt.Println("Title:", response.Data.Title)
	fmt.Println("Content:", response.Data.Content)
	fmt.Println("Completed:", response.Data.Completed == 1)
}
