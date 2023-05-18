package engine

import "github.com/wikylyu/stability/api"

type Engine struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type EngineClient struct {
	c *api.Client
}
