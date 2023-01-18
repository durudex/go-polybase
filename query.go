/*
 * Copyright © 2022 V1def
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

type Query interface {
	Get(ctx context.Context, resp any) error
	Before(cursor string) Query
	After(cursor string) Query
	Limit(num int) Query
	Sort(field string, direction ...string) Query
	Where(field string, op WhereOperator, value any) Query
}

type query struct {
	client   Client
	endpoint string
	param    map[string]any
}

func newQuery(client Client, endpoint string) Query {
	return &query{client: client, endpoint: endpoint, param: make(map[string]any)}
}

func (q *query) Get(ctx context.Context, resp any) error {
	req := &Request{
		Endpoint: q.endpoint + q.build(),
		Method:   "GET",
	}

	return q.client.MakeRequest(ctx, req, resp)
}

func (q *query) build() string {
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

func (q *query) Before(cursor string) Query {
	q.param["before"] = cursor

	return q
}

func (q *query) After(cursor string) Query {
	q.param["after"] = cursor

	return q
}

func (q *query) Limit(num int) Query {
	q.param["limit"] = num

	return q
}

func (q *query) Sort(field string, direction ...string) Query {
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

func (q *query) Where(field string, op WhereOperator, value any) Query {
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
