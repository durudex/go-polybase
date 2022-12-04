/*
 * Copyright © 2022 V1def
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polybase

import "context"

type Option interface {
	Get(ctx context.Context, v any) error
}

type option struct {
	client Client
	req    *Request
}

func NewOption(client Client, req *Request) Option {
	return option{client: client, req: req}
}

func (o option) Get(ctx context.Context, v any) error {
	return o.client.MakeRequest(ctx, o.req, v)
}
