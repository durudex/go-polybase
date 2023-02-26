/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

import (
	"context"
	"encoding/json"
	"fmt"
)

type WhereOperator string

func (o WhereOperator) String() string { return string(o) }

const (
	GreaterThan        WhereOperator = "$gt"
	LessThan           WhereOperator = "$lt"
	GreaterThanOrEqual WhereOperator = "$gte"
	LessThanOrEqual    WhereOperator = "$lte"
	Equal              WhereOperator = "$eq"
)

type Query[T any] interface {
	Get(ctx context.Context) *Response[T]
	Before(cursor string) Query[T]
	After(cursor string) Query[T]
	Limit(num int) Query[T]
	Sort(field string, direction ...string) Query[T]
	Where(field string, op WhereOperator, value any) Query[T]
}

type query[T any] struct {
	client   Client
	endpoint string
	param    map[string]any
}

func newQuery[T any](client Client, endpoint string) Query[T] {
	return &query[T]{client: client, endpoint: endpoint, param: make(map[string]any)}
}

func (q *query[T]) Get(ctx context.Context) *Response[T] {
	req := &Request{
		Endpoint: q.endpoint + q.build(),
		Method:   "GET",
	}

	var resp Response[T]

	q.client.MakeRequest(ctx, req, &resp)

	return &resp
}

func (q *query[T]) build() string {
	res := "?"

	for i, value := range q.param {
		switch value.(type) {
		case string:
			res += fmt.Sprintf("%s=%s", i, value)
		case int:
			res += fmt.Sprintf("%s=%d", i, value)
		default:
			b, err := json.Marshal(value)
			if err != nil {
				return ""
			}

			res += fmt.Sprintf("%s=%s", i, string(b))
		}

		res += "&"
	}

	return res
}

func (q *query[T]) Before(cursor string) Query[T] {
	q.param["before"] = cursor

	return q
}

func (q *query[T]) After(cursor string) Query[T] {
	q.param["after"] = cursor

	return q
}

func (q *query[T]) Limit(num int) Query[T] {
	q.param["limit"] = num

	return q
}

func (q *query[T]) Sort(field string, direction ...string) Query[T] {
	if direction == nil {
		direction = append(direction, "asc")
	}

	if v, ok := q.param["sort"]; !ok {
		q.param["sort"] = [][]string{{field, direction[0]}}
	} else {
		q.param["sort"] = append(v.([][]string), []string{field, direction[0]})
	}

	return q
}

func (q *query[T]) Where(field string, op WhereOperator, value any) Query[T] {
	if _, ok := q.param["where"]; !ok {
		q.param["where"] = make(map[string]any)
	}

	if op == Equal {
		q.param["where"].(map[string]any)[field] = value
	} else {
		if _, ok := q.param["where"].(map[string]any)[field]; !ok {
			q.param["where"].(map[string]any)[field] = make(map[string]any)
		}

		q.param["where"].(map[string]any)[field].(map[string]any)[op.String()] = value
	}

	return q
}
