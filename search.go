package itunes

import "context"

func (c *Client) Search(ctx context.Context, opts ...RequestOption) (*iTunesResult, error) {
	v := processOptions(opts...).urlParams

	itunesUrl := c.baseURL + "search?" + v.Encode()

	var result iTunesResult

	err := c.get(ctx, itunesUrl, &result)

	if err != nil {
		return nil, err
	}

	return &result, err
}
