package awxclient

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Credential struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
}

func (c *Client) CreateCredential(cred Credential) (*Credential, error) {
	resp, err := c.httpClient.R().SetBody(cred).Post("/api/v2/credentials/")
	if err != nil {
		return nil, err
	}

	var createdCred Credential
	err = json.Unmarshal(resp.Body(), &createdCred)
	if err != nil {
		return nil, err
	}

	return &createdCred, nil
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
