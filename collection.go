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

type Collection[T any] interface {
	Query[T]

	Record(id string) RecordDoer[T]
	Create(ctx context.Context, args []any) *SingleResponse[T]
}

type collection[T any] struct {
	name   string
	client Client
}

func NewCollection[T any](client Client, name string) Collection[T] {
	if client.Config().DefaultNamespace != "" {
		name = client.Config().DefaultNamespace + "/" + name
	}

	return &collection[T]{name: url.QueryEscape(name), client: client}
}

func (c *collection[T]) Get(ctx context.Context) *Response[T] {
	req := &Request{
		Endpoint: fmt.Sprintf("/collections/%s/records", c.name),
		Method:   "GET",
	}

	var resp Response[T]

	c.client.MakeRequest(ctx, req, &resp)

	return &resp
}

func (c *collection[T]) Before(cursor string) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf("/collections/%s/records", c.name)).Before(cursor)
}

func (c *collection[T]) After(cursor string) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf("/collections/%s/records", c.name)).After(cursor)
}

func (c *collection[T]) Limit(num int) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf("/collections/%s/records", c.name)).Limit(num)
}

func (c *collection[T]) Sort(field string, direction ...string) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf("/collections/%s/records", c.name)).Sort(field, direction...)
}

func (c collection[T]) Where(field string, op WhereOperator, value any) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf("/collections/%s/records", c.name)).Where(field, op, value)
}

func (c *collection[T]) Record(id string) RecordDoer[T] {
	return newRecordDoer[T](c.client,
		fmt.Sprintf("/collections/%s/records/%s", c.name, url.QueryEscape(id)))
}

func (c *collection[T]) Create(ctx context.Context, args []any) *SingleResponse[T] {
	req := &Request{
		Endpoint: fmt.Sprintf("/collections/%s/records", c.name),
		Method:   "POST",
		Body:     Body{Args: args},
	}

	var resp SingleResponse[T]

	c.client.MakeRequest(ctx, req, &resp)

	return &resp
}
