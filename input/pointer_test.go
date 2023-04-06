/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package input_test

import (
	"reflect"
	"testing"

	"github.com/durudex/go-polybase/input"
)

var ParsePointerTests = map[string]struct {
	arg  any
	want []any
}{
	"String": {
		arg:  new(string),
		want: []any{new(string)},
	},
	"Integer": {
		arg:  new(int),
		want: []any{new(int)},
	},
	"Boolean": {
		arg:  new(bool),
		want: []any{new(bool)},
	},
	"Struct": {
		arg: &ParseStructMock{
			Integer: 1,
			String:  "durudex",
			Boolean: true,
		},
		want: []any{1, "durudex", true},
	},
	"Foreign": {
		arg: &input.Foreign{CollectionID: "example/example", ID: "1"},
		want: []any{
			&input.Foreign{CollectionID: "example/example", ID: "1"},
		},
	},
}

func TestParsePointer(t *testing.T) {
	for name, test := range ParsePointerTests {
		t.Run(name, func(t *testing.T) {
			got := input.ParsePointer(test.arg)

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: want does not match")
			}
		})
	}
}

func BenchmarkParsePointer(b *testing.B) {
	for name, test := range ParsePointerTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input.ParsePointer(test.arg)
			}
		})
	}
}
