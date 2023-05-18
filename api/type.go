package api

type TextPrompt struct {
	Text   string  `json:"text"`
	Weight float64 `json:"weight"`
}
