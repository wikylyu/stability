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

### Text to image generation

```golang
package main

import (
	"fmt"
	"os"

	"github.com/wikylyu/stability"
	"github.com/wikylyu/stability/generation"
)

func main() {
	ai := stability.New(&stability.Config{
		ApiKey:     os.Getenv("STABILITY_API_KEY"),
	})

	textPrompts := []generation.TextPrompt{}
	textPrompts = append(textPrompts, generation.TextPrompt{
		Text:   "Surrounded by his loyal allies, The Mandalorian leads a daring mission across the galactic landscape. His determined presence and poised demeanor exemplify leadership and determination. The painting captures the cooperation and camaraderie of the group.",
		Weight: 0.5,
	})
	resp, err := ai.Generation.Text2Image("stable-diffusion-512-v2-1", &generation.Text2ImageRequest{
		TextPrompts: textPrompts,
		StylePreset: generation.StylePresetCinematic,
		Samples:     1,
	})
	if err != nil {
		panic(err)
	}
	for _, image := range resp.Artifacts {
		fmt.Printf("%s: %s\n", image.FinishReason, image.Base64)
	}
}

```

## API List

- [x] user
  - [x] account
  - [x] balance
- [x] engines
  - [x] list
- [ ] generation
  - [x] text-to-image
  - [ ] image-to-image
  - [ ] image-to-image/upscale
  - [ ] image-to-image/masking
