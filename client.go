package seaweed

import (
	"net/http"
	"os"
	"time"
)

// Client represents a seaweed API client
type Client struct {
	APIKey     string
	HTTPClient *http.Client
	CacheAge   time.Duration
	CacheDir   string
}

// NewClient takes an API key and returns a seaweed API client
func NewClient(APIKey string) *Client {
	dur, _ := time.ParseDuration("5m")

	return &Client{
		APIKey,
		&http.Client{},
		dur,
		os.TempDir(),
	}
}

func logRequests() bool {
	return os.Getenv("SW_LOG") != ""
}

func disableCache() bool {
	return os.Getenv("SW_DISABLE_CACHE") != ""
}
