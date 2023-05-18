package main

import (
	"encoding/base64"
	"image"
	"image/png"
	"os"
	"strings"

	"github.com/wikylyu/stability"
	"github.com/wikylyu/stability/api"
	"github.com/wikylyu/stability/generation"
)

func main() {
	ai := stability.New(&stability.Config{
		ApiKey:     os.Getenv("STABILITY_API_KEY"),
		HttpsProxy: os.Getenv("https_proxy"),
	})

	textPrompts := []api.TextPrompt{}
	textPrompts = append(textPrompts, api.TextPrompt{
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
	for _, img := range resp.Artifacts {
		// fmt.Printf("%s: %s\n", img.FinishReason, img.Base64)
		if err := saveBase64("generated.png", img.Base64); err != nil {
			panic(err)
		}
	}

	f, err := os.Open("generated.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	resp2, err := ai.Generation.Image2Image("stable-diffusion-512-v2-1", &generation.Image2ImageRequest{
		TextPrompts: []api.TextPrompt{
			{Text: "a dragon", Weight: 0.5},
		},
		StylePreset: generation.StylePreset3D_Model,
		InitImage:   f,
	})
	if err != nil {
		panic(err)
	}
	for _, img := range resp2.Artifacts {
		if err := saveBase64("generated_image.png", img.Base64); err != nil {
			panic(err)
		}
	}

}

func saveBase64(filename, b64 string) error {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64))
	i, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, i); err != nil {
		return err
	}
	return nil
}
