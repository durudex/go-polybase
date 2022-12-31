/*
 * Copyright © 2022 Durudex
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

// Record structure stores the Polybase record.
type Record[T any] struct {
	// Block field stores block data from the blockchain.
	Block Block `json:"block"`

	// Data field stores data with the specified type.
	Data T `json:"data"`
}

// RecordDoer interface stores methods for interacting with the specified Polybase record.
type RecordDoer interface {
	// Get method sends a request to getting collection record by the specified ID and decodes
	// the returned value.
	Get(ctx context.Context, resp any) error

	// Call method calls a function from the Polybase collection scheme with the specified
	// arguments. To make it easier to pass arguments, you can pass a structure using the
	// ParseInput function.
	Call(ctx context.Context, fc string, args []any, resp any) error
}

// recordDoer structure implements all methods of the RecordDoer interface.
type recordDoer struct {
	client   Client
	endpoint string
}

// newRecordDoer function returns a new record doer.
func newRecordDoer(client Client, endpoint string) RecordDoer {
	return recordDoer{client: client, endpoint: endpoint}
}

// Get method sends a request to getting collection record by the specified ID and decodes
// the returned value.
func (r recordDoer) Get(ctx context.Context, resp any) error {
	req := &Request{
		Endpoint: r.endpoint,
		Method:   "GET",
	}

	return r.client.MakeRequest(ctx, req, resp)
}

// Call method calls a function from the Polybase collection scheme with the specified arguments.
// To make it easier to pass arguments, you can pass a structure using the ParseInput function.
func (r recordDoer) Call(ctx context.Context, fc string, args []any, resp any) error {
	req := &Request{
		Endpoint: r.endpoint + fmt.Sprintf("/call/%s", url.QueryEscape(fc)),
		Method:   "POST",
		Body:     Body{Args: args},
	}

	return r.client.MakeRequest(ctx, req, resp)
}
