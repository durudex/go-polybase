/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package polybase

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

const (
	// The ClientHeaderKey constant stores the value of the client key for
	// an HTTP header.
	ClientHeaderKey = "X-Polybase-Client"

	// The clientHeaderPrefix constant stores the value of the client prefix
	// for an HTTP header. This value is added to the beginning of the client
	// name specified in the configuration.
	//
	// Example: "durudex/go-polybase:default-name"
	clientHeaderPrefix = "durudex/go-polybase:"
)

// Client interface stores methods for interaction with Polybase Node.
type Client interface {
	// MakeRequest method makes a request with the specified settings and
	// decodes the JSON response.
	MakeRequest(ctx context.Context, req *Request, resp any) error

	Config() *Config
}

// Request structure stores data what used for creating a new HTTP request.
type Request struct {
	Endpoint string
	Method   string
	Body     Body
}

// Body structure stores data what used in client POST HTTP requests.
type Body struct {
	Args []any `json:"args"`
}

// Response structure stores data when returned from multiple responses.
type Response[T any] struct {
	Data   []Record[T] `json:"data,omitempty"`
	Cursor Cursor      `json:"cursor,omitempty"`
	Error  Error       `json:"error,omitempty"`
}

// SingleResponse structure stores data when returned from single response.
type SingleResponse[T any] struct {
	Block Block `json:"block,omitempty"`
	Data  T     `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

// client structure implements all methods of the Client interface.
type client struct {
	cfg  *Config
	doer *http.Client
}

// To start using the go-polybase client, you need to crete a new client
// instance. This can be done using the internal New() function, which
// returns a new instance with either a specified configuration or the
// default configuration.
//
// To create an instance with a specified configuration, you need to pass
// a pointer of Config value as an argument to the New(...) function. This
// can be useful if you want to use specific settings, for example, if you
// have your own configuration file.
//
//	client := polybase.New(&polybase.Config{
//		...
//	}
//
// If you want to use the default configuration, you can simply call the New()
// function without any arguments. The client will be created with the default
// configuration set in the go-polybase module.
//
//	client := polybase.New()
func New(configs ...*Config) Client {
	var cfg *Config

	if configs != nil {
		cfg = configs[0]
	} else {
		cfg = new(Config)
	}

	cfg.configure()

	return &client{cfg: cfg, doer: http.DefaultClient}
}

// MakeRequest method makes a request with the specified settings and decodes
// the JSON response.
func (c *client) MakeRequest(ctx context.Context, req *Request, resp any) error {
	// Creating a new HTTP request.
	rwc, err := c.newRequest(ctx, req)
	if err != nil {
		return err
	}

	// Sending an HTTP request and returns a HTTP response.
	re, err := c.doer.Do(rwc)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	// Decoding the JSON returned response.
	return json.NewDecoder(re.Body).Decode(resp)
}

// newRequest method returns a new HTTP request by client Request and configuration.
func (c *client) newRequest(ctx context.Context, req *Request) (*http.Request, error) {
	var body io.Reader

	// Creating a new body reader when body arguments not empty.
	if req.Body.Args != nil {
		b, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(b)
	}

	url := c.cfg.URL + req.Endpoint

	// Creating a new HTTP request with user context.
	rwc, err := http.NewRequestWithContext(ctx, req.Method, url, body)
	if err != nil {
		return nil, err
	}

	name := clientHeaderPrefix + c.cfg.Name

	// Addition HTTP headers for the request.
	rwc.Header.Add(ClientHeaderKey, name)

	return rwc, nil
}

func (c *client) Config() *Config { return c.cfg }
