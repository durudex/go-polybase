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
	client   Client
	endpoint string
}

func NewOption(client Client, endpoint string) Option {
	return option{client: client, endpoint: endpoint}
}

func (o option) Get(ctx context.Context, v any) error {
	req := &Request{
		Endpoint: o.endpoint,
		Method:   "GET",
	}

	return o.client.MakeRequest(ctx, req, v)
}
