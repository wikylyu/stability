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
		HttpsProxy: os.Getenv("https_proxy"),
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
