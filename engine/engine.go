package engine

import (
	"encoding/json"

	"github.com/wikylyu/stability/api"
)

func NewClient(c *api.Client) *EngineClient {
	return &EngineClient{c: c}
}

/*
 * List all engines available to your organization/user
 */
func (c *EngineClient) List() ([]*Engine, error) {
	data, err := c.c.DoJson("GET", "/v1/engines/list", nil)
	if err != nil {
		return nil, err
	}
	resp := make([]*Engine, 0)
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
