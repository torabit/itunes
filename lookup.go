package itunes

import (
	"context"
)

func (c *Client) Lookup(ctx context.Context, lookupOpt LookupOption, opts ...RequestOption) (*iTunesResult, error) {
	v := processOptions(opts...)
	lookupOpt(&v)

	itunesUrl := c.baseURL + "lookup?" + v.urlParams.Encode()

	var result iTunesResult

	err := c.get(ctx, itunesUrl, &result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
