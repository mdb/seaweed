package seaweed

import (
	"net/http"
	"os"
)

type Client struct {
	ApiKey     string
	HttpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey,
		&http.Client{},
	}
}

func LogRequests() bool {
	return os.Getenv("SW_LOG") != ""
}
