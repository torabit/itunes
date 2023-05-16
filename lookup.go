package itunes

import (
	"context"
	"net/url"
)

func (c *Client) Lookup(ctx context.Context, lookupOpt LookupOption, opts ...RequestOption) (*iTunesResult, error) {
	v := processOptions(opts...).urlParams

	k := requestOption{
		urlParams: url.Values{},
	}

	lookupOpt(&k)

	itunesUrl := c.baseURL + "lookup?" + k.urlParams.Encode() + v.Encode()

	var result iTunesResult

	err := c.get(ctx, itunesUrl, &result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
