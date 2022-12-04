package polybase

import (
	"fmt"
	"net/url"
)

type Collection interface {
	Record(id string) Option
}

type collection struct {
	name   string
	client Client
}

func NewCollection(name string, client Client) Collection {
	return &collection{name: url.QueryEscape(name), client: client}
}

func (c *collection) Record(id string) Option {
	uri := URL + fmt.Sprintf("/collections/%s/records/%s", c.name, url.QueryEscape(id))

	req := &Request{
		URL:    uri,
		Method: "GET",
	}

	return NewOption(c.client, req)
}
