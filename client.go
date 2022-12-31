/*
 * Copyright © 2022 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Client interface {
	MakeRequest(ctx context.Context, req *Request, resp any) error
}

type Request struct {
	Endpoint string
	Method   string
	Body     Body
}

type Body struct {
	Args []any `json:"args"`
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

type client struct {
	url  string
	doer *http.Client
}

func NewClient(url string) Client {
	// TODO: add custom client
	return &client{url: url, doer: http.DefaultClient}
}

func (c *client) MakeRequest(ctx context.Context, req *Request, resp any) error {
	var body io.Reader

	if req.Body.Args != nil {
		b, err := json.Marshal(req.Body)
		if err != nil {
			return err
		}

		body = bytes.NewReader(b)
	}

	rc, err := http.NewRequestWithContext(ctx, req.Method, c.url+req.Endpoint, body)
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
