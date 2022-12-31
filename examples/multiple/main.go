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

	"github.com/durudex/go-polybase"
)

type Collection struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

func main() {
	db := polybase.New(polybase.Config{URL: polybase.TestnetURL})
	coll := db.Collection("polybase/todo")

	var response polybase.Response[Collection]

	coll.Get(context.Background(), &response)

	fmt.Println("After Cursor", response.Cursor.After)
	fmt.Println("Before Cursor", response.Cursor.Before)

	for _, data := range response.Data {
		fmt.Println("Block Hash:", data.Block.Hash)
		fmt.Println("ID:", data.Data.ID)
		fmt.Println("Code:", data.Data.Code)
	}
}
