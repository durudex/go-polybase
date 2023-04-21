/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package polybase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/durudex/go-polybase/input"
)

const (
	recordsEndpointFormat = "/collections/%s/records"
	recordEndpointFormat  = recordsEndpointFormat + "/%s"
)

// RecordDoer interface stores methods for interacting with the specified
// Polybase record.
type Record[T any] interface {
	// Get method sends a request to getting collection record by the
	// specified ID and decodes the returned value.
	Get(ctx context.Context) *Response[T]

	// Call method calls a function from the Polybase collection scheme
	// with the specified arguments.
	Call(ctx context.Context, fc string, args ...any) *Response[T]
}

// record structure implements all methods of the Record interface.
type record[T any] struct {
	client   Client
	endpoint string
	config   *input.Foreign
}

// newRecordDoer function returns a new record doer.
func newRecord[T any](client Client, name, id string) Record[T] {
	endpoint := fmt.Sprintf(recordEndpointFormat, name, url.QueryEscape(id))
	config := &input.Foreign{CollectionID: name, ID: id}

	return &record[T]{
		client:   client,
		endpoint: endpoint,
		config:   config,
	}
}

// Get method sends a request to getting collection record by the specified
// ID and decodes the returned value.
func (r *record[T]) Get(ctx context.Context) *Response[T] {
	defer recoverFunc(ctx, r.client.Config().RecoverHandler)

	req := &Request{
		Endpoint: r.endpoint,
		Method:   http.MethodGet,
	}

	var resp Response[T]

	if err := r.client.MakeRequest(ctx, req, &resp); err != nil {
		panic("error: getting record: " + err.Error())
	}

	return &resp
}

// Call method calls a function from the Polybase collection scheme with
// the specified arguments.
func (r *record[T]) Call(ctx context.Context, fc string, args ...any) *Response[T] {
	defer recoverFunc(ctx, r.client.Config().RecoverHandler)

	req := &Request{
		Endpoint: r.endpoint + fmt.Sprintf("/call/%s", url.QueryEscape(fc)),
		Method:   http.MethodPost,
		Body:     Body{Args: input.Parse(args)},
	}

	var resp Response[T]

	if err := r.client.MakeRequest(ctx, req, &resp); err != nil {
		panic("error: call collection function: " + err.Error())
	}

	return &resp
}

func (r *record[T]) Reference() *input.Foreign { return r.config }
