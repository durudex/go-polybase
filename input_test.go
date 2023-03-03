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

func TestParseInput(t *testing.T) {
	type Mock struct {
		String  string
		Integer int
		Boolean bool
	}

	tests := []struct {
		name string
		args []any
		want []any
	}{
		{
			name: "Pointer Struct",
			args: []any{&Mock{String: "string", Integer: 1, Boolean: true}},
			want: []any{"string", 1, true},
		},
		{
			name: "Struct",
			args: []any{Mock{String: "string", Integer: 1, Boolean: true}},
			want: []any{"string", 1, true},
		},
		{
			name: "Any Slice",
			args: []any{[]any{"string", 1, true}},
			want: []any{"string", 1, true},
		},
		{
			name: "String Slice",
			args: []any{[]string{"1", "2", "3"}},
			want: []any{"1", "2", "3"},
		},
		{
			name: "Multiple",
			args: []any{"string", 1, true},
			want: []any{"string", 1, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := polybase.ParseInput(tt.args)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: args does not match")
			}
		})
	}
}
