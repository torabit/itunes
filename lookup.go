package itunes

import (
	"context"
	"fmt"
)

func (c *Client) Lookup(ctx context.Context, lookupOpt LookupOption, opts ...RequestOption) (*iTunesResult, error) {
	v := processOptions(opts...)
	lookupOpt(&v)

	itunesUrl := c.baseURL + "lookup?" + v.urlParams.Encode()

	fmt.Printf("%v\n", itunesUrl)

	var result iTunesResult

	err := c.get(ctx, itunesUrl, &result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
