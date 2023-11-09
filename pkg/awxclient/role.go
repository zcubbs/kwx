package awxclient

import "encoding/json"

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ListRoles retrieves all roles.
func (c *Client) ListRoles() ([]Role, error) {
	resp, err := c.httpClient.R().Get("/api/v2/roles/")
	if err != nil {
		return nil, err
	}
	var roles []Role
	err = json.Unmarshal(resp.Body(), &roles)
	return roles, err
}
