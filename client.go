package seaweed

import "net/http"

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
