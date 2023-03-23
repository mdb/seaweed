package seaweed

import (
	"net/http"
	"time"

	logging "github.com/op/go-logging"
)

// Clock is a clock interface used to report the current time.
type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time {
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
		realClock{},
	}
}

// Forecast fetches the full, multi-day forecast for a given spot.
func (c *Client) Forecast(spot string) ([]Forecast, error) {
	forecasts := []Forecast{}
	err := getForecast(c, spot, &forecasts)
	if err != nil {
		return forecasts, err
	}

	return forecasts, nil
}

// Today fetches the today's forecast for a given spot.
func (c *Client) Today(spot string) ([]Forecast, error) {
	today := c.clock.Now().UTC().Day()
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchDays(forecasts, today), nil
}

// Tomorrow fetches tomorrow's forecast for a given spot.
func (c *Client) Tomorrow(spot string) ([]Forecast, error) {
	tomorrowDate := c.clock.Now().UTC().Day() + 1
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchDays(forecasts, tomorrowDate), nil
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
