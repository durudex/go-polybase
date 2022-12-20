/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"context"
	"time"

	"github.com/v1def/go-polybase"
)

func main() {
	coll := polybase.New(polybase.Config{
		URL: polybase.TestnetURL,
	}).Collection("polybase/todo")

	coll.Create(context.Background(), []any{time.Now().String(), "go-polybase", "example", 0})
}
