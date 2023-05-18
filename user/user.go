package user

import (
	"encoding/json"

	"github.com/wikylyu/stability/api"
)

func NewClient(c *api.Client) *UserClient {
	return &UserClient{c: c}
}

/*
 * Get Account information
 */
func (c *UserClient) GetAccount() (*User, error) {
	data, err := c.c.DoJson("GET", "/v1/user/account", nil)
	if err != nil {
		return nil, err
	}
	var resp User
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
