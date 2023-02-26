/*
 * Copyright © 2022-2023 Durudex
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
type RecordDoer[T any] interface {
	// Get method sends a request to getting collection record by the specified ID and decodes
	// the returned value.
	Get(ctx context.Context) *SingleResponse[T]

	// Call method calls a function from the Polybase collection scheme with the specified
	// arguments. To make it easier to pass arguments, you can pass a structure using the
	// ParseInput function.
	Call(ctx context.Context, fc string, args []any) *SingleResponse[T]
}

// recordDoer structure implements all methods of the RecordDoer interface.
type recordDoer[T any] struct {
	client   Client
	endpoint string
}

// newRecordDoer function returns a new record doer.
func newRecordDoer[T any](client Client, endpoint string) RecordDoer[T] {
	return recordDoer[T]{client: client, endpoint: endpoint}
}

// Get method sends a request to getting collection record by the specified ID and decodes
// the returned value.
func (r recordDoer[T]) Get(ctx context.Context) *SingleResponse[T] {
	req := &Request{
		Endpoint: r.endpoint,
		Method:   "GET",
	}

	var resp SingleResponse[T]

	if err := r.client.MakeRequest(ctx, req, &resp); err != nil {
		panic("error getting record: " + err.Error())
	}

	return &resp
}

// Call method calls a function from the Polybase collection scheme with the specified arguments.
// To make it easier to pass arguments, you can pass a structure using the ParseInput function.
func (r recordDoer[T]) Call(ctx context.Context, fc string, args []any) *SingleResponse[T] {
	req := &Request{
		Endpoint: r.endpoint + fmt.Sprintf("/call/%s", url.QueryEscape(fc)),
		Method:   "POST",
		Body:     Body{Args: args},
	}

	var resp SingleResponse[T]

	if err := r.client.MakeRequest(ctx, req, &resp); err != nil {
		panic("error call collection function: " + err.Error())
	}

	return &resp
}
