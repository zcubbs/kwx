package awxclient

import "github.com/go-resty/resty/v2"

func (c *Client) Ping() (*resty.Response, error) {
	resp, err := c.httpClient.R().Get("/api/v2/ping/")
	return resp, err
}
