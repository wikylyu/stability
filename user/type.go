package user

import "github.com/wikylyu/stability/api"

type Organization struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	IsDefault bool   `json:"is_default"`
}

type User struct {
	ID             string         `json:"id"`
	Email          string         `json:"email"`
	ProfilePicture string         `json:"profile_picture"`
	Organizations  []Organization `json:"organizations"`
}

type Blance struct {
	Credits float64 `json:"credits"`
}

type UserClient struct {
	c *api.Client
}
