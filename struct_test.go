/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase_test

import (
	"reflect"
	"testing"

	"github.com/durudex/go-polybase"
)

func TestParseInput(t *testing.T) {
	type Test struct {
		String  string
		Integer int
		Boolean bool
	}

	want := []any{"string", 1, true}

	got := polybase.ParseInput(&Test{String: "string", Integer: 1, Boolean: true})

	if !reflect.DeepEqual(want, got) {
		t.Fatal("error: want does not match")
	}
}
