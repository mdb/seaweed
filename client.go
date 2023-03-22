package seaweed

import (
	"net/http"
	"os"
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
	CacheAge   time.Duration
	CacheDir   string
	Log        *logging.Logger
	clock      Clock
}

// NewClient takes an API key and returns a seaweed API client
func NewClient(APIKey string) *Client {
	dur, _ := time.ParseDuration("5m")

	return &Client{
		APIKey,
		&http.Client{},
		dur,
		os.TempDir(),
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
	today := c.clock.Now().Day()
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchDays(forecasts, today), nil
}

// Tomorrow fetches tomorrow's forecast for a given spot.
func (c *Client) Tomorrow(spot string) ([]Forecast, error) {
	tomorrowDate := c.clock.Now().Day() + 1
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchDays(forecasts, tomorrowDate), nil
}

// Weekend fetches the weekend's forecast for a given spot.
func (c *Client) Weekend(spot string) ([]Forecast, error) {
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchWeekendDays(forecasts), nil
}

func disableCache() bool {
	return os.Getenv("SW_DISABLE_CACHE") != ""
}
