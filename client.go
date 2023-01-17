/*
 * Copyright © 2022-2023 Durudex
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
	cfg  Config
	doer *http.Client
}

func NewClient(cfg Config) Client {
	return &client{cfg: cfg, doer: http.DefaultClient}
}

func (c *client) MakeRequest(ctx context.Context, req *Request, resp any) error {
	rwc, err := c.newRequest(ctx, req)
	if err != nil {
		return err
	}

	re, err := c.doer.Do(rwc)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	return json.NewDecoder(re.Body).Decode(resp)
}

func (c *client) newRequest(ctx context.Context, req *Request) (*http.Request, error) {
	var body io.Reader

	if req.Body.Args != nil {
		b, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(b)
	}

	url := c.cfg.URL + req.Endpoint

	rwc, err := http.NewRequestWithContext(ctx, req.Method, url, body)
	if err != nil {
		return nil, err
	}

	name := "durudex/go-polybase:" + c.cfg.Name
	rwc.Header.Add("X-Polybase-Client", name)

	return rwc, nil
}
