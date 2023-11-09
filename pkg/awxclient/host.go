package awxclient

import (
	"encoding/json"
	"fmt"
)

type Host struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	InventoryID int    `json:"inventory"`
}

// CreateHost creates a new host within a specified inventory.
func (c *Client) CreateHost(host Host) (*Host, error) {
	resp, err := c.httpClient.R().SetBody(host).Post("/api/v2/hosts/")
	if err != nil {
		return nil, err
	}
	var result Host
	err = json.Unmarshal(resp.Body(), &result)
	return &result, err
}

// GetHost retrieves details of a host by its ID.
func (c *Client) GetHost(id int) (*Host, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("/api/v2/hosts/%d/", id))
	if err != nil {
		return nil, err
	}
	var host Host
	err = json.Unmarshal(resp.Body(), &host)
	return &host, err
}
