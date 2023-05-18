package generation

import (
	"encoding/json"
	"fmt"

	"github.com/wikylyu/stability/api"
)

func NewClient(c *api.Client) *GenerationClient {
	return &GenerationClient{c: c}
}

/*
 * Generate a new image from a text prompt
 */
func (c *GenerationClient) Text2Image(engineID string, req *Text2ImageRequest) (*Text2ImageResponse, error) {
	data, err := c.c.DoJson("POST", fmt.Sprintf("/v1/generation/%s/text-to-image", engineID), req)
	if err != nil {
		return nil, err
	}
	var resp Text2ImageResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *GenerationClient) Image2Image(engineID string, req *Image2ImageRequest) (*Image2ImageResponse, error) {
	data, err := c.c.DoForm("POST", fmt.Sprintf("/v1/generation/%s/image-to-image", engineID), req)
	if err != nil {
		return nil, err
	}
	var resp Image2ImageResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
