/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

import (
	"fmt"
	"net/url"
)

type Query interface {
	Before(string) Query
	After(string) Query
	Params() []string
}

type query struct{ params []string }

func NewQuery() Query { return query{} }

func (q query) Before(s string) Query {
	q.params = append(q.params, fmt.Sprintf("before=%s", url.QueryEscape(s)))

	return q
}

func (q query) After(s string) Query {
	q.params = append(q.params, fmt.Sprintf("after=%s", url.QueryEscape(s)))

	return q
}

func (q query) Params() []string { return q.params }
