# Stability

Unofficial Go implemention for [Stability AI REST API](https://platform.stability.ai/rest-api).

## Usage

### Get account information

```golang
package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/stability"
)

func main() {
	ai := stability.New(&stability.Config{
		ApiKey:     os.Getenv("STABILITY_API_KEY"),
	})
	user, err := ai.User.GetAccount()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %s\n", user.ID)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Profile Pictures: %s\n", user.ProfilePicture)
	for _, org := range user.Organizations {
		fmt.Printf("\t Org Name: %s\n", org.Name)
	}
}

```
