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

type ParseStructMock struct {
	Integer int
	String  string
	Boolean bool
}

type ParseStructPointerMock struct {
	Integer *int
	String  *string
	Boolean *bool
}

var ParseStructTests = map[string]struct {
	arg  any
	want []any
}{
	"Default": {
		arg: ParseStructMock{
			Integer: 1,
			String:  "hello",
			Boolean: true,
		},
		want: []any{1, "hello", true},
	},
	"Pointer": {
		arg: ParseStructPointerMock{
			Integer: new(int),
			String:  new(string),
			Boolean: new(bool),
		},
		want: []any{0, "", false},
	},
	"Foreign": {
		arg: struct{ F input.Foreign }{
			input.Foreign{
				CollectionID: "example/example",
				ID:           "1",
			},
		},
		want: []any{input.Foreign{
			CollectionID: "example/example",
			ID:           "1",
		}},
	},
}

func TestParseStruct(t *testing.T) {
	for name, test := range ParseStructTests {
		t.Run(name, func(t *testing.T) {
			got := input.ParseStruct(test.arg)

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: want does not match")
			}
		})
	}
}

func BenchmarkParseStruct(b *testing.B) {
	for name, test := range ParseStructTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input.ParseStruct(test.arg)
			}
		})
	}
}
