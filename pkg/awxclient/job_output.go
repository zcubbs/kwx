package awxclient

import "fmt"

type JobOutput struct {
	Stdout string `json:"stdout"`
}

func (c *Client) GetJobOutput(jobID int) (JobOutput, error) {
	var output JobOutput
	endpoint := fmt.Sprintf("/api/v2/jobs/%d/stdout/", jobID)
	resp, err := c.httpClient.R().SetResult(&output).Get(endpoint)
	if err != nil {
		return JobOutput{}, err
	}
	if resp.IsError() {
		return JobOutput{}, fmt.Errorf("API responded with status code %d", resp.StatusCode())
	}
	return output, nil
}
