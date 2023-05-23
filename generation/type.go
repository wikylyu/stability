package generation

import (
	"encoding/base64"
	"image"
	"os"
	"strings"

	"github.com/wikylyu/stability/api"
)

const (
	ClipGuidancePresetNONE       = "NONE"
	ClipGuidancePresetFAST_BLUE  = "FAST_BLUE"
	ClipGuidancePresetFAST_GREEN = "FAST_GREEN"
	ClipGuidancePresetSIMPLE     = "SIMPLE"
	ClipGuidancePresetSLOW       = "SLOW"
	ClipGuidancePresetSLOWER     = "SLOWER"
	ClipGuidancePresetSLOWEST    = "SLOWEST"
)

const (
	SamplerDDIM                 = "DDIM"
	SamplerDDPM                 = "DDPM"
	SamplerK_DPMPP_2M           = "K_DPMPP_2M"
	SamplerK_DPMPP_2S_ANCESTRAL = "K_DPMPP_2S_ANCESTRAL"
	SamplerK_DPM_2              = "K_DPM_2"
	SamplerK_DPM_2_ANCESTRAL    = "K_DPM_2_ANCESTRAL"
	SamplerK_EULER              = "K_EULER"
	SamplerK_EULER_ANCESTRAL    = "K_EULER_ANCESTRAL"
	SamplerK_HEUN               = "K_HEUN"
	SamplerK_LMS                = "K_LMS"
)

const (
	StylePreset3D_Model          = "3d-model"
	StylePresetAnalog_Film       = "analog-film"
	StylePresetAnime             = "anime"
	StylePresetCinematic         = "cinematic"
	StylePresetComic_Book        = "comic-book"
	StylePresetDigital_Art       = "digital-art"
	StylePresetEnhance           = "enhance"
	StylePresetFantasy_Art       = "fantasy-art"
	StylePresetIsometric         = "iosmetric"
	StylePresetLine_Art          = "line-art"
	StylePresetLow_Poly          = "low-poly"
	StylePresetModeling_Compound = "modeling-compound"
	StylePresetNeon_Punk         = "neon-punk"
	StylePresetOrigami           = "origami"
	StylePresetPhotographic      = "photographic"
	StylePresetPixel_Art         = "pixel-art"
	StylePresetTile_Texture      = "tile-texture"
)

// type TextPrompt struct {
// 	Text   string  `json:"text"`
// 	Weight float64 `json:"weight"`
// }

type Text2ImageRequest struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`

	TextPrompts []api.TextPrompt `json:"text_prompts"`

	CfgScale           float64 `json:"cfg_scale,omitempty"`
	ClipGuidancePreset string  `json:"clip_guidance_preset,omitempty"`
	Sampler            string  `json:"sampler,omitempty"`
	Samples            int     `json:"samples,omitempty"`
	Seed               int64   `json:"seed,omitempty"`
	Steps              int     `json:"steps,omitempty"`
	StylePreset        string  `json:"style_preset,omitempty"`

	Extras interface{} `json:"extras,omitempty"`
}

type Image struct {
	Base64       string `json:"base64"`
	FinishReason string `json:"finishReason"`
	Seed         int64  `json:"seed"`
}

func (img *Image) ToImage() (image.Image, string, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(img.Base64))
	return image.Decode(reader)
}

type Text2ImageResponse struct {
	Artifacts []Image `json:"artifacts"`
}

type GenerationClient struct {
	c *api.Client
}

type Image2ImageRequest struct {
	TextPrompts        []api.TextPrompt `form:"text_prompts"`
	InitImage          *os.File         `form:"init_image"`
	InitImageMode      string           `form:"init_image_mode"`
	ImageStrength      float64          `form:"image_strength"`
	CfgScale           float64          `form:"cfg_scale"`
	ClipGuidancePreset string           `form:"clip_guidance_preset"`
	Sampler            string           `form:"sampler"`
	Samples            int              `form:"samples"`
	Seed               int64            `form:"seed"`
	Steps              int              `form:"steps"`
	StylePreset        string           `form:"style_preset"`

	Extras interface{} `form:"extras"`
}

type Image2ImageResponse struct {
	Artifacts []Image `json:"artifacts"`
}
