/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package codegen_test

import (
	"reflect"
	"testing"

	"github.com/v1def/go-polybase/codegen"
)

func TestNewConfig(t *testing.T) {
	want := &codegen.Config{
		Collections: []string{"durudex/user", "durudex/post"},
		Package:     "generated",
		Directory:   "generated",
	}

	got, err := codegen.NewConfig("fixtures/config.yml")
	if err != nil {
		t.Fatal("error: creating a new config: ", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatal("error: config does not match")
	}
}
