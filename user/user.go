package user

import (
	"encoding/json"

	"github.com/wikylyu/stability/api"
)

func NewClient(c *api.Client) *UserClient {
	return &UserClient{c: c}
}

/*
 * Get information about the account associated with the provided API key
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

/*
 * Get the credit balance of the account/organization associated with the API key
 */
func (c *UserClient) GetBalance() (*Blance, error) {
	data, err := c.c.DoJson("GET", "/v1/user/balance", nil)
	if err != nil {
		return nil, err
	}
	var resp Blance
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
