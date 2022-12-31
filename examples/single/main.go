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
	coll := polybase.New(polybase.Config{
		URL: polybase.TestnetURL,
	}).Collection("Collection")

	var response polybase.SingleResponse[Collection]

	coll.Record("Collection").Get(context.Background(), &response)

	fmt.Println("Block Hash:", response.Block.Hash)
	fmt.Println("ID:", response.Data.ID)
	fmt.Println("Code:", response.Data.Code)
}
