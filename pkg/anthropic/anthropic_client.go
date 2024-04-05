package anthropic

import "net/http"

// Client represents the Anthropic API client and its configuration.
type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

type ClientOption func(client *Client) error

func New(apiKey string, baseURL string, httpClient *http.Client) (*Client, error) {
	c := &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: httpClient,
	}

	return c, nil
}
