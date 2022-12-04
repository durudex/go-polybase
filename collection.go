package polybase

import (
	"context"
	"fmt"
	"net/url"
)

type Collection interface {
	Record(id string) Option
	Get(ctx context.Context, resp any) error
}

type collection struct {
	name   string
	client Client
}

func NewCollection(name string, client Client) Collection {
	return &collection{name: url.QueryEscape(name), client: client}
}

func (c *collection) Record(id string) Option {
	req := &Request{
		Endpoint: fmt.Sprintf("/collections/%s/records/%s", c.name, url.QueryEscape(id)),
		Method:   "GET",
	}

	return NewOption(c.client, req)
}

func (c *collection) Get(ctx context.Context, resp any) error {
	req := &Request{
		Endpoint: fmt.Sprintf("/collections/%s/records", c.name),
		Method:   "GET",
	}

	return c.client.MakeRequest(ctx, req, resp)
}
