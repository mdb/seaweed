package seaweed

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	logging "github.com/op/go-logging"
)

// Clock is a clock interface used to report the current time.
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
	Log        *logging.Logger
	clock      Clock
}

// NewClient takes an API key and returns a seaweed API client
func NewClient(APIKey string) *Client {
	return &Client{
		APIKey,
		&http.Client{},
		NewLogger(logging.INFO),
		RealClock{},
	}
}

// Forecast fetches the full, multi-day forecast for a given spot.
func (c *Client) Forecast(spot string) ([]Forecast, error) {
	forecasts, err := c.getForecast(spot)
	if err != nil {
		return forecasts, err
	}

	return forecasts, nil
}

// Today fetches the today's forecast for a given spot.
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

// Tomorrow fetches tomorrow's forecast for a given spot.
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

// Weekend fetches the weekend's forecast for a given spot.
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

	err = json.Unmarshal(body, &forecasts)
	if err != nil {
		return forecasts, err
	}

	return forecasts, nil
}

func (c *Client) get(url string) ([]byte, error) {
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%s returned HTTP status code %d", url, resp.StatusCode)
	}

	c.Log.Debugf("url=%s http_status=%d response_body=%s", url, resp.StatusCode, string(bodyContents))

	return bodyContents, err
}
