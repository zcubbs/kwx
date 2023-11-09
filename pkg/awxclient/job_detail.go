package awxclient

import "fmt"

type JobDetail struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (c *Client) GetJobDetail(jobID int) (JobDetail, error) {
	var job JobDetail
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/", jobID)
	resp, err := c.httpClient.R().SetResult(&job).Get(endpoint)
	if err != nil {
		return JobDetail{}, err
	}
	if resp.IsError() {
		return JobDetail{}, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}
	return job, nil
}
