package seaweed

import (
	"net/http"
	"os"
	"time"
)

type Client struct {
	ApiKey     string
	HttpClient *http.Client
	CacheAge   time.Duration
	CacheDir   string
}

func NewClient(apiKey string) *Client {
	dur, _ := time.ParseDuration("5m")

	return &Client{
		apiKey,
		&http.Client{},
		dur,
		os.TempDir(),
	}
}

func LogRequests() bool {
	return os.Getenv("SW_LOG") != ""
}

func DisableCache() bool {
	return os.Getenv("SW_DISABLE_CACHE") != ""
}
