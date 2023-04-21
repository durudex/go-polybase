/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package polybase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/durudex/go-polybase/input"
)

type Collection[T any] interface {
	Query[T]

	Record(id string) Record[T]
	Create(ctx context.Context, args ...any) *Response[T]
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

func (c *collection[T]) Get(ctx context.Context) *ResponseList[T] {
	defer recoverFunc(ctx, c.client.Config().RecoverHandler)

	req := &Request{
		Endpoint: fmt.Sprintf(recordsEndpointFormat, c.name),
		Method:   http.MethodGet,
	}

	var resp ResponseList[T]

	if err := c.client.MakeRequest(ctx, req, &resp); err != nil {
		panic("error: getting collection records: " + err.Error())
	}

	return &resp
}

func (c *collection[T]) Before(cursor string) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf(recordsEndpointFormat, c.name)).Before(cursor)
}

func (c *collection[T]) After(cursor string) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf(recordsEndpointFormat, c.name)).After(cursor)
}

func (c *collection[T]) Limit(num int) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf(recordsEndpointFormat, c.name)).Limit(num)
}

func (c *collection[T]) Sort(field string, direction ...string) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf(recordsEndpointFormat, c.name)).Sort(field, direction...)
}

func (c *collection[T]) Where(field string, op WhereOperator, value any) Query[T] {
	return newQuery[T](c.client,
		fmt.Sprintf(recordsEndpointFormat, c.name)).Where(field, op, value)
}

func (c *collection[T]) Record(id string) Record[T] {
	return newRecord[T](c.client, c.name, id)
}

func (c *collection[T]) Create(ctx context.Context, args ...any) *Response[T] {
	defer recoverFunc(ctx, c.client.Config().RecoverHandler)

	req := &Request{
		Endpoint: fmt.Sprintf(recordsEndpointFormat, c.name),
		Method:   http.MethodPost,
		Body:     Body{Args: input.Parse(args)},
	}

	var resp Response[T]

	if err := c.client.MakeRequest(ctx, req, &resp); err != nil {
		panic("error: creating a new record instance: " + err.Error())
	}

	return &resp
}
