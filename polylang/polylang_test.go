/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polylang_test

import (
	"os"
	"testing"

	"github.com/v1def/go-polybase/polylang"

	"github.com/alecthomas/participle/v2"
)

func TestCollection(t *testing.T) {
	parser := participle.MustBuild[polylang.Collection]()

	f, err := os.Open("./fixtures/collection.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	if got.Name != "Polybase" {
		t.Fatal("error: collection name does not match")
	}
}

func TestField(t *testing.T) {
	parser := participle.MustBuild[polylang.Collection]()

	f, err := os.Open("./fixtures/field.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	switch {
	case got.Items[0].Field.Name != "polybase":
		t.Fatal("error: field name does not match")
	case got.Items[0].Field.Optional != true:
		t.Fatal("error: field optional does not match")
	case got.Items[0].Field.Type != "string":
		t.Fatal("error: field type does not match")
	}
}

func TestFunction(t *testing.T) {
	parser := participle.MustBuild[polylang.Collection]()

	f, err := os.Open("./fixtures/function.polylang")
	if err != nil {
		t.Fatal("error: opening fixtures file: ", err)
	}
	defer f.Close()

	got, err := parser.Parse("", f)
	if err != nil {
		t.Fatal("error: parsing polylang file: ", err)
	}

	switch {
	case got.Items[0].Function.Name != "constructor":
		t.Fatal("error: function name does not match")
	case got.Items[0].Function.Parameters[0].Name != "id":
		t.Fatal("error: parameter name does not match")
	case got.Items[0].Function.Parameters[0].Type != "string":
		t.Fatal("error: parameter type does not match")
	case got.Items[0].Function.Parameters[1].Name != "name":
		t.Fatal("error: parameter name does not match")
	case got.Items[0].Function.Parameters[1].Optional != true:
		t.Fatal("error: parameter optional does not match")
	}
}
