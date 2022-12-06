/*
 * Copyright © 2022 V1def
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
	Record(id string) Option
	Get(ctx context.Context, v any, params ...string) error
	Create(ctx context.Context, v []any) error
}

type collection struct {
	name   string
	client Client
}

func NewCollection(name string, client Client) Collection {
	return &collection{name: url.QueryEscape(name), client: client}
}

func (c *collection) Record(id string) Option {
	return NewOption(c.client, fmt.Sprintf("/collections/%s/records/%s", c.name, url.QueryEscape(id)))
}

func (c *collection) Get(ctx context.Context, v any, params ...string) error {
	req := &Request{
		Endpoint: fmt.Sprintf("/collections/%s/records", c.name) + buildParam(params),
		Method:   "GET",
	}

	return c.client.MakeRequest(ctx, req, v)
}

func (c *collection) Create(ctx context.Context, v []any) error {
	req := &Request{
		Endpoint: fmt.Sprintf("/collections/%s/records", c.name),
		Method:   "POST",
		Body:     Body{Args: v},
	}

	return c.client.MakeRequest(ctx, req, v)
}
