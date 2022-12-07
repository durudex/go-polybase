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
	case got.Items[0].Field.Required != true:
		t.Fatal("error: field required does not match")
	case got.Items[0].Field.Type != "string":
		t.Fatal("error: field type does not match")
	}
}
