/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package input_test

import (
	"reflect"
	"testing"

	"github.com/durudex/go-polybase/input"
)

var ParseMapTests = map[string]struct{ arg, want any }{
	"Integer Key": {
		arg:  map[int]int{1: 2, 3: 4, 5: 6, 7: 8},
		want: map[int]int{1: 2, 3: 4, 5: 6, 7: 8},
	},
	"String Key": {
		arg:  map[string]int{"1": 1, "2": 2, "3": 3},
		want: map[string]int{"1": 1, "2": 2, "3": 3},
	},
	"Foreign": {
		arg: map[int]input.Foreign{
			1: {CollectionID: "2", ID: "3"},
			4: {CollectionID: "5", ID: "6"},
		},
		want: map[int]input.Foreign{
			1: {CollectionID: "2", ID: "3"},
			4: {CollectionID: "5", ID: "6"},
		},
	},
	"Custom Foreign": {
		arg: map[int]ParseForeignMock{
			1: {ID: "2"}, 3: {ID: "4"},
		},
		want: map[int]*input.Foreign{
			1: {CollectionID: "example/example", ID: "2"},
			3: {CollectionID: "example/example", ID: "4"},
		},
	},
}

func TestParseMap(t *testing.T) {
	for name, test := range ParseMapTests {
		t.Run(name, func(t *testing.T) {
			got := input.ParseMap(test.arg)

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: want does not match")
			}
		})
	}
}

func BenchmarkParseMap(b *testing.B) {
	for name, test := range ParseMapTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input.ParsePointer(test.arg)
			}
		})
	}
}
