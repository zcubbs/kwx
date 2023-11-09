package awxclient

import (
	"encoding/json"
	"fmt"
)

type NotificationTemplate struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// GetNotificationTemplate retrieves a notification template by its ID.
func (c *Client) GetNotificationTemplate(id int) (*NotificationTemplate, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("/api/v2/notification_templates/%d/", id))
	if err != nil {
		return nil, err
	}
	var notification NotificationTemplate
	err = json.Unmarshal(resp.Body(), &notification)
	return &notification, err
}
