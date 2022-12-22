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

// Option interface stores methods for interacting with the specified Polybase record.
type Option interface {
	// Get method sends a request to retrieve all fields of the collection by the specified ID and
	// decodes the returned value.
	Get(ctx context.Context, resp any) error

	// Call method calls a function from the Polybase collection scheme with the specified arguments.
	// To make it easier to pass arguments, you can pass a structure using the ParseInput function.
	Call(ctx context.Context, fc string, args []any, resp any) error
}

// option structure implements all methods of the Option interface.
type option struct {
	client   Client
	endpoint string
}

// NewOption function returns a new record option.
func NewOption(client Client, endpoint string) Option {
	return option{client: client, endpoint: endpoint}
}

// Get method sends a request to retrieve all fields of the collection by the specified ID and
// decodes the returned value.
func (o option) Get(ctx context.Context, resp any) error {
	req := &Request{
		Endpoint: o.endpoint,
		Method:   "GET",
	}

	return o.client.MakeRequest(ctx, req, resp)
}

// Call method calls a function from the Polybase collection scheme with the specified arguments.
// To make it easier to pass arguments, you can pass a structure using the ParseInput function.
func (o option) Call(ctx context.Context, fc string, args []any, resp any) error {
	req := &Request{
		Endpoint: o.endpoint + fmt.Sprintf("/call/%s", url.QueryEscape(fc)),
		Method:   "POST",
		Body:     Body{Args: args},
	}

	return o.client.MakeRequest(ctx, req, resp)
}
