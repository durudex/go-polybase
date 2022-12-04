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
	return o.client.Record(ctx, o.req, v)
}
