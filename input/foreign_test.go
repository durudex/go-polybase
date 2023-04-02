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

type ParseForeignMock struct{ ID string }

func (m ParseForeignMock) Reference() *input.Foreign {
	return &input.Foreign{
		CollectionID: "example/example",
		ID:           m.ID,
	}
}

var ParseForeignTests = map[string]struct{ arg, want any }{
	"Default": {
		arg: input.Foreign{
			CollectionID: "example/example",
			ID:           "1",
		},
		want: input.Foreign{
			CollectionID: "example/example",
			ID:           "1",
		},
	},
	"Pointer": {
		arg: &input.Foreign{
			CollectionID: "example/example",
			ID:           "1",
		},
		want: &input.Foreign{
			CollectionID: "example/example",
			ID:           "1",
		},
	},
	"Custom": {
		arg: ParseForeignMock{ID: "1"},
		want: &input.Foreign{
			CollectionID: "example/example",
			ID:           "1",
		},
	},
	"Custom Pointer": {
		arg: &ParseForeignMock{ID: "1"},
		want: &input.Foreign{
			CollectionID: "example/example",
			ID:           "1",
		},
	},
}

func TestParseForeign(t *testing.T) {
	for name, test := range ParseForeignTests {
		t.Run(name, func(t *testing.T) {
			got := input.ParseForeign(test.arg)

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: want does not match")
			}
		})
	}
}

func BenchmarkParseForeign(b *testing.B) {
	for name, test := range ParseForeignTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input.ParseForeign(test.arg)
			}
		})
	}
}
