package stability

import (
	"time"

	"github.com/wikylyu/stability/engine"
	"github.com/wikylyu/stability/generation"
	"github.com/wikylyu/stability/user"
)

type Session struct {
	User       *user.UserClient
	Engine     *engine.EngineClient
	Generation *generation.GenerationClient
}

type Config struct {
	BaseURL         string // endpoint, optional
	ApiKey          string // api key, required
	MaxIdleConns    int
	IdleConnTimeout time.Duration
	HttpsProxy      string
}
