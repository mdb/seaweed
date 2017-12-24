package seaweed

import (
	"net/http"
	"os"
	"time"

	logging "github.com/op/go-logging"
)

// Client represents a seaweed API client
type Client struct {
	APIKey     string
	HTTPClient *http.Client
	CacheAge   time.Duration
	CacheDir   string
	Log        *logging.Logger
}

// NewClient takes an API key and returns a seaweed API client
func NewClient(APIKey string) *Client {
	dur, _ := time.ParseDuration("5m")

	return &Client{
		APIKey,
		&http.Client{},
		dur,
		os.TempDir(),
		NewLogger(),
	}
}

func disableCache() bool {
	return os.Getenv("SW_DISABLE_CACHE") != ""
}
