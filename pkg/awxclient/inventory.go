package awxclient

import (
	"encoding/json"
	"fmt"
)

type Inventory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateInventory creates a new inventory.
func (c *Client) CreateInventory(inv Inventory) (*Inventory, error) {
	resp, err := c.httpClient.R().SetBody(inv).Post("/api/v2/inventories/")
	if err != nil {
		return nil, err
	}
	var result Inventory
	err = json.Unmarshal(resp.Body(), &result)
	return &result, err
}

// GetInventory retrieves an inventory by its ID.
func (c *Client) GetInventory(id int) (*Inventory, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("/api/v2/inventories/%d/", id))
	if err != nil {
		return nil, err
	}
	var inv Inventory
	err = json.Unmarshal(resp.Body(), &inv)
	return &inv, err
}
