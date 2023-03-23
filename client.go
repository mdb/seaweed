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

// Clock is a clock interface used to report the current time.
// It exists largely to test the Client#Tomorrow method.
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
	APIKey     string
	HTTPClient *http.Client
	Logger     *logrus.Logger
	clock      Clock
}

// NewClient takes an API key and returns a seaweed API client
func NewClient(APIKey string) *Client {
	return &Client{
		APIKey,
		&http.Client{},
		logrus.New(),
		RealClock{},
	}
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
	today := []Forecast{}
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
	tomorrow := []Forecast{}
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
	weekendFs := []Forecast{}
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
	url := fmt.Sprintf("http://magicseaweed.com/api/%s/forecast/?spot_id=%s", c.APIKey, spotID)
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
			return forecasts, err
		}

		return forecasts, errors.New(errResp.ErrorResponse.ErrorMsg)
	default:
		err = json.Unmarshal(body, &forecasts)
		if err != nil {
			return forecasts, err
		}

		return forecasts, nil
	}
}

func (c *Client) get(url string) ([]byte, error) {
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sanitizedURL := strings.Replace(url, c.APIKey, "<REDACTED>", 1)

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%s returned HTTP status code %d", sanitizedURL, resp.StatusCode)
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
