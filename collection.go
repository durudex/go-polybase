/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

import (
	"context"
	"fmt"
	"net/url"
)

type Collection interface {
	Query

	Record(id string) RecordDoer
	Create(ctx context.Context, args []any, resp any) error
}

type collection struct {
	name   string
	client Client
}

func newCollection(name string, client Client) Collection {
	return &collection{name: url.QueryEscape(name), client: client}
}

func (c *collection) Get(ctx context.Context, resp any) error {
	req := &Request{
		Endpoint: fmt.Sprintf("/collections/%s/records", c.name),
		Method:   "GET",
	}

	return c.client.MakeRequest(ctx, req, resp)
}

func (c *collection) Before(cursor string) Query {
	return newQuery(c.client, fmt.Sprintf("/collections/%s/records", c.name)).Before(cursor)
}

func (c *collection) After(cursor string) Query {
	return newQuery(c.client, fmt.Sprintf("/collections/%s/records", c.name)).After(cursor)
}

func (c *collection) Limit(num int) Query {
	return newQuery(c.client, fmt.Sprintf("/collections/%s/records", c.name)).Limit(num)
}

func (c *collection) Sort(field string, direction ...string) Query {
	return newQuery(c.client,
		fmt.Sprintf("/collections/%s/records", c.name)).Sort(field, direction...)
}

func (c collection) Where(field string, op WhereOperator, value any) Query {
	return newQuery(c.client,
		fmt.Sprintf("/collections/%s/records", c.name)).Where(field, op, value)
}

func (c *collection) Record(id string) RecordDoer {
	return newRecordDoer(c.client,
		fmt.Sprintf("/collections/%s/records/%s", c.name, url.QueryEscape(id)))
}

func (c *collection) Create(ctx context.Context, args []any, resp any) error {
	req := &Request{
		Endpoint: fmt.Sprintf("/collections/%s/records", c.name),
		Method:   "POST",
		Body:     Body{Args: args},
	}

	return c.client.MakeRequest(ctx, req, resp)
}
