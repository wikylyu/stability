package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/stability"
)

func main() {
	ai := stability.New(&stability.Config{
		ApiKey:     os.Getenv("STABILITY_API_KEY"),
		HttpsProxy: os.Getenv("https_proxy"),
	})

	engines, err := ai.Engine.List()
	if err != nil {
		panic(err)
	}
	for _, engine := range engines {
		fmt.Printf("%s: %s(%s)\n", engine.ID, engine.Name, engine.Type)
	}
}
