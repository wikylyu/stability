package generation

import "github.com/wikylyu/stability/api"

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

type TextPrompt struct {
	Text   string  `json:"text"`
	Weight float64 `json:"weight"`
}

type Text2ImageRequest struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`

	TextPrompts []TextPrompt `json:"text_prompts"`

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

type Text2ImageResponse struct {
	Artifacts []Image `json:"artifacts"`
}

type GenerationClient struct {
	c *api.Client
}
