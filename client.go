package polybase

import (
	"context"
	"encoding/json"
	"net/http"
)

type Client interface {
	MakeRequest(ctx context.Context, req *Request, resp any) error
}

type Request struct {
	URL    string
	Method string
}

type Response[T any] struct {
	Data   []Record[T] `json:"data,omitempty"`
	Cursor Cursor      `json:"cursor,omitempty"`
	Error  Error       `json:"error,omitempty"`
}

type SingleResponse[T any] struct {
	Block Block `json:"block,omitempty"`
	Data  T     `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

type client struct{ doer *http.Client }

func NewClient() Client {
	// TODO: add custom client
	return &client{doer: http.DefaultClient}
}

func (c *client) MakeRequest(ctx context.Context, req *Request, resp any) error {
	rc, err := http.NewRequestWithContext(ctx, req.Method, req.URL, nil)
	if err != nil {
		return err
	}

	re, err := c.doer.Do(rc)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	return json.NewDecoder(re.Body).Decode(resp)
}
