/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"context"

	"github.com/v1def/go-polybase"
)

func main() {
	coll := polybase.New(polybase.Config{
		URL: polybase.TestnetURL,
	}).Collection("polybase/todo")

	coll.Record("1").Call(context.Background(), "update", []any{"1", "go-polybase", "example", 0})
}
