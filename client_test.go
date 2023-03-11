/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase_test

import (
	"testing"

	"github.com/durudex/go-polybase"
)

var NewTests = map[string]struct {
	configs []*polybase.Config
	want    *polybase.Config
}{
	"OK": {
		configs: []*polybase.Config{
			{
				URL:              "126.0.0.1:80",
				Name:             "example-bot",
				DefaultNamespace: "example",
			},
		},
		want: &polybase.Config{
			URL:              "126.0.0.1:80",
			Name:             "example-bot",
			DefaultNamespace: "example",
		},
	},
	"Empty": {
		want: &polybase.Config{
			URL:  polybase.DefaultURL,
			Name: polybase.DefaultName,
		},
	},
}

func TestNew(t *testing.T) {
	for name, test := range NewTests {
		t.Run(name, func(t *testing.T) {
			got := polybase.New(test.configs...).Config()

			switch {
			case got.URL != test.want.URL:
				t.Fatal("error: URL does not match")
			case got.Name != test.want.Name:
				t.Fatal("error: name does not match")
			case got.DefaultNamespace != test.want.DefaultNamespace:
				t.Fatal("error: default namespace does not match")
			}
		})
	}
}
