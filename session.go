package stability

import (
	"github.com/wikylyu/stability/api"
	"github.com/wikylyu/stability/engine"
	"github.com/wikylyu/stability/user"
)

func New(cfg *Config) *Session {
	client := api.NewClient(cfg.BaseURL, cfg.ApiKey, cfg.MaxIdleConns, cfg.IdleConnTimeout, cfg.HttpsProxy)
	return &Session{
		User:   user.NewClient(client),
		Engine: engine.NewClient(client),
	}
}
