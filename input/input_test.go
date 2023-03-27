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

var ParseTests = map[string]struct{ args, want []any }{
	"Default": {
		args: []any{1, 2, 3, "hello", true},
		want: []any{1, 2, 3, "hello", true},
	},
	"Slice": {
		args: []any{[]any{1, "hello"}, []int{1, 2, 3}},
		want: []any{1, "hello", []int{1, 2, 3}},
	},
	"Struct": {
		args: []any{
			ParseStructMock{Integer: 1, String: "hello", Boolean: true},
			ParseStructMock{Integer: 2, String: "world", Boolean: false},
		},
		want: []any{1, "hello", true, 2, "world", false},
	},
	"Foreign": {
		args: []any{
			input.Foreign{CollectionID: "example/example", ID: "1"},
		},
		want: []any{
			input.Foreign{CollectionID: "example/example", ID: "1"},
		},
	},
	"Pointer": {
		args: []any{
			&ParseStructMock{Integer: 1, String: "durudex", Boolean: true},
			&input.Foreign{CollectionID: "example/example", ID: "1"},
			new(string), new(int), new(bool),
		},
		want: []any{
			1, "durudex", true, &input.Foreign{
				CollectionID: "example/example",
				ID:           "1",
			}, "", 0, false,
		},
	},
}

func TestParse(t *testing.T) {
	for name, test := range ParseTests {
		t.Run(name, func(t *testing.T) {
			got := input.Parse(test.args)

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: want does not match")
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	for name, test := range ParseTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input.Parse(test.args)
			}
		})
	}
}
