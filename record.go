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
	"net/http"
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
	// arguments.
	Call(ctx context.Context, fc string, args ...any) *SingleResponse[T]
}

// recordDoer structure implements all methods of the RecordDoer interface.
type recordDoer[T any] struct {
	client   Client
	endpoint string
}

// newRecordDoer function returns a new record doer.
func newRecordDoer[T any](client Client, endpoint string) RecordDoer[T] {
	return &recordDoer[T]{client: client, endpoint: endpoint}
}

// Get method sends a request to getting collection record by the specified ID and decodes
// the returned value.
func (r *recordDoer[T]) Get(ctx context.Context) *SingleResponse[T] {
	defer recoverFunc(ctx, r.client.Config().RecoverHandler)

	req := &Request{
		Endpoint: r.endpoint,
		Method:   http.MethodGet,
	}

	var resp SingleResponse[T]

	if err := r.client.MakeRequest(ctx, req, &resp); err != nil {
		panic("error getting record: " + err.Error())
	}

	return &resp
}

// Call method calls a function from the Polybase collection scheme with the specified arguments.
func (r *recordDoer[T]) Call(ctx context.Context, fc string, args ...any) *SingleResponse[T] {
	defer recoverFunc(ctx, r.client.Config().RecoverHandler)

	req := &Request{
		Endpoint: r.endpoint + fmt.Sprintf("/call/%s", url.QueryEscape(fc)),
		Method:   http.MethodPost,
		Body:     Body{Args: ParseInput(args)},
	}

	var resp SingleResponse[T]

	if err := r.client.MakeRequest(ctx, req, &resp); err != nil {
		panic("error call collection function: " + err.Error())
	}

	return &resp
}
