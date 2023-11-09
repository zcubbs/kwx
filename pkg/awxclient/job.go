package awxclient

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Job struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (c *Client) LaunchJob(templateID int) (*Job, error) {
	resp, err := c.httpClient.R().Post(fmt.Sprintf("/api/v2/job_templates/%d/launch/", templateID))
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.Unmarshal(resp.Body(), &job)
	return &job, err
}

// GetJob retrieves details of a job by its ID.
func (c *Client) GetJob(id int) (*Job, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("/api/v2/jobs/%d/", id))
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.Unmarshal(resp.Body(), &job)
	return &job, err
}

func (c *Client) CancelJob(jobID int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/cancel/", jobID)
	resp, err := c.httpClient.R().Post(endpoint)
	return resp, err
}
