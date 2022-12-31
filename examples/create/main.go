/*
 * Copyright © 2022 Durudex
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
	Completed bool   `json:"completed"`
}

func main() {
	db := polybase.New(polybase.Config{URL: polybase.TestnetURL})
	coll := db.Collection("polybase/todo")

	var response polybase.SingleResponse[Todo]

	coll.Create(context.Background(), []any{time.Now().String(), "go-polybase", "example", 0}, &response)

	fmt.Println("ID:", response.Data.ID)
	fmt.Println("Title:", response.Data.Title)
	fmt.Println("Content:", response.Data.Content)
	fmt.Println("Completed:", response.Data.Completed)
}
