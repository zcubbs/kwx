package awxclient

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Credential struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Client) AddCredential(cred Credential) (*resty.Response, error) {
	resp, err := c.httpClient.R().SetBody(cred).Post("/api/v2/credentials/")
	return resp, err
}

func (c *Client) DeleteCredential(credID int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/api/v2/credentials/%d/", credID)
	resp, err := c.httpClient.R().Delete(endpoint)
	return resp, err
}

func (c *Client) UpdateCredential(credID int, updatedCred Credential) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/api/v2/credentials/%d/", credID)
	resp, err := c.httpClient.R().SetBody(updatedCred).Put(endpoint)
	return resp, err
}
