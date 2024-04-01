package anthropic

import (
	"fmt"
	"net/http"
)

// Client represents the Anthropic API client and its configuration.
//type Client struct {
//	httpClient *http.Client
//	apiKey     string
//	baseURL    string
//}

// NewClient initializes a new Anthropic API client with the required headers.
func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, ErrAnthropicApiKeyRequired
	}
	options := &options{
		apiKey:     apiKey,
		baseURL:    BASE_URL,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(options)
	}

	client, err := New(options.apiKey, options.baseURL, options.httpClient)
	fmt.Printf("client %#v \n", client)
	return client, err
}

//func NewClient1(apiKey string) (*Client, error) {
//	if apiKey == "" {
//		return nil, ErrAnthropicApiKeyRequired
//	}
//
//	client := &Client{
//		httpClient: &http.Client{},
//		apiKey:     apiKey,
//		baseURL:    "https://api.anthropic.com",
//	}
//
//	return client, nil
//}
