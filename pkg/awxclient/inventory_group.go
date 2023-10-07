package awxclient

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	InventoryID int    `json:"inventory"`
}

// CreateGroup creates a new group within a specified inventory.
func (c *Client) CreateGroup(group Group) (*Group, error) {
	resp, err := c.httpClient.R().SetBody(group).Post("/api/v2/groups/")
	if err != nil {
		return nil, err
	}
	var result Group
	err = json.Unmarshal(resp.Body(), &result)
	return &result, err
}

// GetGroup retrieves details of a group by its ID.
func (c *Client) GetGroup(id int) (*Group, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("/api/v2/groups/%d/", id))
	if err != nil {
		return nil, err
	}
	var group Group
	err = json.Unmarshal(resp.Body(), &group)
	return &group, err
}
