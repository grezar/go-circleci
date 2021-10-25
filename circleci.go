package circleci

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-querystring/query"
)

const (
	userAgent = "go-circleci"

	DefaultAddress  = "https://circleci.com"
	DefaultBasePath = "/api/v2/"
)

type Config struct {
	Address    string
	BasePath   string
	Token      string
	Headers    http.Header
	HTTPClient *http.Client
}

func DefaultConfig() *Config {
	config := &Config{
		Address:    DefaultAddress,
		BasePath:   DefaultBasePath,
		Token:      os.Getenv("CIRCLECI_TOKEN"),
		Headers:    make(http.Header),
		HTTPClient: &http.Client{},
	}

	config.Headers.Set("User-Agent", userAgent)

	return config
}

type Client struct {
	baseURL *url.URL
	token   string
	headers http.Header
	http    *http.Client

	Contexts Contexts
}

func NewClient(cfg *Config) (*Client, error) {
	config := DefaultConfig()

	if cfg != nil {
		if cfg.Address != "" {
			config.Address = cfg.Address
		}
		if cfg.BasePath != "" {
			config.BasePath = cfg.BasePath
		}
		if cfg.Token != "" {
			config.Token = cfg.Token
		}
		for k, v := range cfg.Headers {
			config.Headers[k] = v
		}
		if cfg.HTTPClient != nil {
			config.HTTPClient = cfg.HTTPClient
		}
	}

	baseURL, err := url.Parse(config.Address)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %v", err)
	}

	if config.Token == "" {
		return nil, fmt.Errorf("API token is required")
	}

	client := &Client{
		baseURL: baseURL,
		token:   config.Token,
		headers: config.Headers,
		http:    config.HTTPClient,
	}

	client.Contexts = &contexts{client: client}

	return client, nil
}

func (c *Client) newRequest(method string, path string, v interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	reqHeaders := make(http.Header)
	reqHeaders.Set("Circle-Token", c.token)
	reqHeaders.Set("Accept", "application/vnd.api+json")

	var body interface{}
	switch method {
	case "GET":
		if v != nil {
			q, err := query.Values(v)
			if err != nil {
				return nil, err
			}
			u.RawQuery = q.Encode()
		}
	case "DELETE", "PATCH", "POST", "PUT":
		reqHeaders.Set("Content-Type", "application/vnd.api+json")
		body = v
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	for k, v := range c.headers {
		req.Header[k] = v
	}

	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) error {
	reqWithCtx := req.WithContext(ctx)

	resp, err := c.http.Do(reqWithCtx)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return err
		}
	}
	defer resp.Body.Close()

	if err := checkResponseCode(resp); err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil
		}
		if decErr != nil {
			err = decErr
		}
	}

	return err
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func checkResponseCode(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode <= 299 {
		return nil
	}

	switch r.StatusCode {
	case 401:
		return ErrUnauthorized
	case 404:
		return ErrNotFound
	}

	var errResponse ErrorResponse
	err := json.NewDecoder(r.Body).Decode(&errResponse)
	if err != nil || errResponse.Message == "" {
		return errors.New(r.Status)
	}

	return errors.New(errResponse.Message)
}
