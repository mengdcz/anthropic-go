package anthropic

import "net/http"

const BASE_URL = "https://api.anthropic.com/v1"

//const BASE_URL = "https://api.anthropic.com"

var (
	ANTHROPIC_API_KEY = "ANTHROPIC_API_KEY"
	//ANTHROPIC_BASE_URL = getEnv("ANTHROPIC_BASE_URL", DEFAULT_ANTHROPIC_BASE_URL)
	//ANTHROPIC_API_KEY  = getEnv("ANTHROPIC_API_KEY", "")
)

type options struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

type Option func(*options)

//	func WithApiKey(apiKey string) Option {
//		return func(o *options) {
//			o.apiKey = apiKey
//		}
//	}
func WithBaseURL(baseURL string) Option {
	return func(o *options) {
		o.baseURL = baseURL
	}
}
