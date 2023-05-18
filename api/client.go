package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

type Client struct {
	baseurl string
	apikey  string

	httpClient *http.Client
}

const DefaultBaseURL = "https://api.stability.ai"

func NewClient(baseurl, apikey string, maxIdleConns int, idleConnTimeout time.Duration, httpsProxy string) *Client {

	if baseurl == "" {
		baseurl = DefaultBaseURL
	}
	if maxIdleConns <= 0 {
		maxIdleConns = 10
	}
	if idleConnTimeout <= 0 {
		idleConnTimeout = 30 * time.Second
	}

	tr := &http.Transport{
		MaxIdleConns:    maxIdleConns,
		IdleConnTimeout: idleConnTimeout,
	}

	/* setup http proxy */
	if httpsProxy != "" {
		proxyUrl, _ := url.Parse(httpsProxy)
		if proxyUrl != nil {
			tr.Proxy = http.ProxyURL(proxyUrl)
		}
	}
	c := &Client{
		baseurl:    baseurl,
		apikey:     apikey,
		httpClient: &http.Client{Transport: tr},
	}

	return c
}

func (c *Client) getJsonRequest(method, path string, body interface{}) (*http.Request, error) {
	fullurl, err := url.JoinPath(c.baseurl, path)
	if err != nil {
		return nil, err
	}
	var reader io.Reader = nil
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(data)
	}
	req, err := http.NewRequest(method, fullurl, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apikey))
	return req, nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var e ErrorResponse
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
		return nil, &e
	}

	return data, nil
}

func (c *Client) Download(path string) (io.ReadCloser, error) {
	req, err := c.getJsonRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var e ErrorResponse
		if err := json.Unmarshal(data, &e); err != nil {
			return nil, err
		}
		return nil, &e
	}
	return resp.Body, nil
}

func (c *Client) DoJson(method, path string, body interface{}) ([]byte, error) {
	req, err := c.getJsonRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *Client) getFormRequest(method, path string, body interface{}) (*http.Request, error) {
	fullurl, err := url.JoinPath(c.baseurl, path)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	if body != nil {
		v := reflect.ValueOf(body).Elem()
		typeOfv := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			key := typeOfv.Field(i).Name
			val := field.Interface()
			if typeOfv.Field(i).Tag.Get("form") != "" {
				key = typeOfv.Field(i).Tag.Get("form")
			}
			switch value := val.(type) {
			case int:
				if value > 0 {
					writer.WriteField(key, fmt.Sprintf("%d", value))
				}
			case int64:
				if value > 0 {
					writer.WriteField(key, fmt.Sprintf("%d", value))
				}
			case float64:
				if value > 0 {
					writer.WriteField(key, fmt.Sprintf("%f", value))
				}
			case string:
				if value != "" {
					writer.WriteField(key, value)
				}
			case *os.File:
				if value != nil {
					part, err := writer.CreateFormFile(key, filepath.Base(value.Name()))
					if err != nil {
						return nil, err
					}
					io.Copy(part, value)
				}
			}

		}
		writer.Close()
	}
	req, err := http.NewRequest(method, fullurl, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apikey))
	return req, nil
}

func (c *Client) DoForm(method, path string, body interface{}) ([]byte, error) {
	req, err := c.getFormRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}
