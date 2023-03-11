/*
 * Copyright © 2022-2023 Durudex
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

type ParseInputMock struct {
	String  string
	Integer int
	Boolean bool
}

var ParseInputTests = map[string]struct {
	args []any
	want []any
}{
	"Pointer Struct": {
		args: []any{&ParseInputMock{"Hello Durudex", 1, true}},
		want: []any{"Hello Durudex", 1, true},
	},
	"Struct": {
		args: []any{ParseInputMock{"string", 1, true}},
		want: []any{"string", 1, true},
	},
	"Struct Unsupported Type": {
		args: []any{struct{ F float32 }{0.1}},
		want: nil,
	},
	"Any Slice": {
		args: []any{[]any{"string", 1, true}},
		want: []any{"string", 1, true},
	},
	"String Slice": {
		args: []any{[]string{"1", "2", "3"}},
		want: []any{"1", "2", "3"},
	},
	"Multiple": {
		args: []any{"string", 1, true},
		want: []any{"string", 1, true},
	},
}

func TestParseInput(t *testing.T) {
	for name, test := range ParseInputTests {
		t.Run(name, func(t *testing.T) {
			got := polybase.ParseInput(test.args)

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: args does not match")
			}
		})
	}
}

func BenchmarkParseInput(b *testing.B) {
	for name, test := range ParseInputTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				polybase.ParseInput(test.args)
			}
		})
	}
}
