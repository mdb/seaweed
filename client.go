// Package seaweed provides a Magic Seaweed API client.
package seaweed

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Clock is a clock interface used to report the current time such that the
// Client#Today and Client#Tomorrow methods can return the proper forecasts
// relative to the current time.
// It exists largely for testing purposes.
type Clock interface {
	Now() time.Time
}

// RealClock implements a Clock.
type RealClock struct{}

// Now returns the current time.Time.
func (RealClock) Now() time.Time {
	return time.Now()
}

// Client represents a seaweed API client
type Client struct {
	// baseURL is the targeted Magic Seaweed API base URL, such as https://magicseaweed.com.
	baseURL string
	// apiKey is a Magic Seaweed API key.
	apiKey string
	// httpClient is a *http.Client.
	httpClient *http.Client
	// Logger is a *logrus.Logger.
	Logger *logrus.Logger
	// clock is a seaweed.Clock used to report the current time/date such that the
	// Client#Tomorrow and Client#Today methods can return the proper forecasts
	// relative to the current time.
	clock Clock
}

// ClientOption configures one or more Client fields.
type ClientOption = func(c *Client)

// WithBaseURL is a ClientOption to configure a *Client's baseURL.
func WithBaseURL(u string) ClientOption {
	return func(c *Client) {
		c.baseURL = u
	}
}

// WithHTTPClient is a ClientOption to configure a *Client's httpClient.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithLogger is a ClientOption to configure a *Client's Logger.
func WithLogger(l *logrus.Logger) ClientOption {
	return func(c *Client) {
		c.Logger = l
	}
}

// WithLogger is a ClientOption to configure a *Client's clock.
func WithClock(clock Clock) ClientOption {
	return func(c *Client) {
		c.clock = clock
	}
}

// NewClient takes an API key and returns a seaweed API client.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		"https://magicseaweed.com",
		apiKey,
		&http.Client{},
		logrus.New(),
		RealClock{},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Forecast fetches the full, multi-day forecast for a given spot ID.
//
// A spot's ID appears in its URL. For example, Ocean City, NJ's spot ID is 391:
// https://magicseaweed.com/Ocean-City-NJ-Surf-Report/391/
//
// Note that the Magic Seaweed API may respond with an HTTP status code of 200
// and a response body reporting an error (see APIError). Forecast attempts to
// handle such instances by returning an error surfacing the response body error
// message.
func (c *Client) Forecast(spot string) ([]Forecast, error) {
	forecasts, err := c.getForecast(spot)
	if err != nil {
		return forecasts, err
	}

	return forecasts, nil
}

// Today fetches the today's forecast for a given spot ID.
func (c *Client) Today(spot string) ([]Forecast, error) {
	var today []Forecast
	now := c.clock.Now().UTC()
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return today, err
	}

	for _, each := range forecasts {
		if each.IsDay(now) {
			today = append(today, each)
		}
	}

	return today, nil
}

// Tomorrow fetches tomorrow's forecast for a given spot ID.
func (c *Client) Tomorrow(spot string) ([]Forecast, error) {
	var tomorrow []Forecast
	tomorrowD := c.clock.Now().UTC().AddDate(0, 0, 1)
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return tomorrow, err
	}

	for _, each := range forecasts {
		if each.IsDay(tomorrowD) {
			tomorrow = append(tomorrow, each)
		}
	}

	return tomorrow, nil
}

// Weekend fetches the weekend's forecast for a given spot ID.
func (c *Client) Weekend(spot string) ([]Forecast, error) {
	var weekendFs []Forecast
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return weekendFs, err
	}

	for _, each := range forecasts {
		if each.IsWeekend() {
			weekendFs = append(weekendFs, each)
		}
	}

	return weekendFs, nil
}

func (c *Client) getForecast(spotID string) ([]Forecast, error) {
	url := fmt.Sprintf("%s/api/%s/forecast/?spot_id=%s", c.baseURL, c.apiKey, spotID)
	forecasts := []Forecast{}
	body, err := c.get(url)
	if err != nil {
		return forecasts, err
	}

	switch {
	case strings.Contains(string(body), "error_response"):
		var errResp APIError
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return forecasts, fmt.Errorf("unexpected API response '%s': %w", body, err)
		}

		return forecasts, errors.New(errResp.ErrorResponse.ErrorMsg)
	default:
		err = json.Unmarshal(body, &forecasts)
		if err != nil {
			return forecasts, fmt.Errorf("unexpected API response '%s': %w", body, err)
		}

		return forecasts, nil
	}
}

// Get is a convenience function that fetches the []Forecast associated with the
// location it's passed using a default Client.
func Get(key, location string) ([]Forecast, error) {
	c := NewClient(key)

	return c.Forecast(location)
}

func (c *Client) get(url string) ([]byte, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sanitizedURL := strings.Replace(url, c.apiKey, "<REDACTED>", 1)
	sanitizedURL = strings.Replace(sanitizedURL, c.baseURL, "", 1)

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("GET %s returned HTTP status code %d", sanitizedURL, resp.StatusCode)
	}

	l := c.Logger.WithFields(
		logrus.Fields{
			"url":         sanitizedURL,
			"http_status": resp.StatusCode,
			"body":        string(body),
		})

	l.Debugf("Magic Seaweed API response")

	return body, err
}
