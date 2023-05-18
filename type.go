package stability

import (
	"time"

	"github.com/wikylyu/stability/user"
)

type Session struct {
	User *user.UserClient
}

type Config struct {
	BaseURL         string // endpoint, optional
	ApiKey          string // api key, required
	MaxIdleConns    int
	IdleConnTimeout time.Duration
	HttpsProxy      string
}
