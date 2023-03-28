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

var ParseIterableTests = map[string]struct {
	arg  any
	want []any
}{
	"Any Slice": {
		arg:  []any{1, "string", true},
		want: []any{1, "string", true},
	},
	"String Slice": {
		arg:  []string{"one", "two", "three"},
		want: []any{[]string{"one", "two", "three"}},
	},
	"Int Slice": {
		arg:  []int{1, 2, 3, 4, 5},
		want: []any{[]int{1, 2, 3, 4, 5}},
	},
	"Bool Slice": {
		arg:  []bool{true, false},
		want: []any{[]bool{true, false}},
	},
	"Foreign Slice": {
		arg: []input.Foreign{
			{CollectionID: "1", ID: "2"}, {CollectionID: "3", ID: "4"},
			{CollectionID: "5", ID: "6"}, {CollectionID: "7", ID: "8"},
		},
		want: []any{[]any{
			input.Foreign{CollectionID: "1", ID: "2"},
			input.Foreign{CollectionID: "3", ID: "4"},
			input.Foreign{CollectionID: "5", ID: "6"},
			input.Foreign{CollectionID: "7", ID: "8"},
		}},
	},
}

func TestParseIterable(t *testing.T) {
	for name, test := range ParseIterableTests {
		t.Run(name, func(t *testing.T) {
			got := input.ParseIterable(test.arg)

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: want does not match")
			}
		})
	}
}

func BenchmarkParseIterable(b *testing.B) {
	for name, test := range ParseIterableTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input.ParseIterable(test.arg)
			}
		})
	}
}
