package seaweed

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	logging "github.com/op/go-logging"
)

func TestNewClient(t *testing.T) {
	client := NewClient("fakeKey")
	age, _ := time.ParseDuration("5m")

	if client.APIKey != "fakeKey" {
		t.Error("NewClient should properly set the API key")
	}
	if client.CacheAge != age {
		t.Error("NewClient should properly set the default 5m cache age")
	}
	if client.CacheDir != os.TempDir() {
		t.Error("NewClient should properly set the default cache directory")
	}
}

func testTools(code int, body string) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, body)
	}))

	tr := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: tr}
	dur, _ := time.ParseDuration("0s")

	client := &Client{
		"fakeKey",
		httpClient,
		dur,
		os.TempDir(),
		NewLogger(logging.INFO),
	}

	return server, client
}

func TestForecast(t *testing.T) {
	server, c := testTools(200, resp)
	defer server.Close()
	forecasts, _ := c.Forecast("123")
	forecast := forecasts[0]

	if forecast.Timestamp != 1443592800 {
		t.Error("Forecast should properly return a Timestamp")
	}

	if forecast.LocalTimestamp != 1443571200 {
		t.Error("Forecast should properly return a LocalTimestamp")
	}

	if forecast.IssueTimestamp != 1443592800 {
		t.Error("Forecast should properly return an IssueTimestamp")
	}

	if forecast.FadedRating != 3 {
		t.Error("Forecast should properly return a FadedRating")
	}

	if forecast.SolidRating != 0 {
		t.Error("Forecast should properly return SolidRating")
	}

	if forecast.Swell.MinBreakingHeight != 5 {
		t.Error("Forecast should properly return Swell.MinBreakingHeight")
	}

	if forecast.Swell.AbsMinBreakingHeight != 4.88 {
		t.Error("Forecast should properly return Swell.AbsMinBreakingHeight")
	}

	if forecast.Swell.Unit != "ft" {
		t.Error("Forecast should properly return Swell.Unit")
	}

	if forecast.Swell.MaxBreakingHeight != 8 {
		t.Error("Forecast should properly return Swell.MaxBreakingHeight")
	}

	if forecast.Swell.AbsMaxBreakingHeight != 7.63 {
		t.Error("Forecast should properly return Swell.AbsMaxBreakingHeight")
	}

	if forecast.Swell.Components.Combined.Height != 7.5 {
		t.Error("Forecast should properly return Swell.Components.Combined.Height")
	}

	if forecast.Swell.Components.Primary.Height != 7.5 {
		t.Error("Forecast should properly return Swell.Components.Primary.Height")
	}

	if forecast.Wind.Speed != 13 {
		t.Error("Forecast should properly return Wind.Speed")
	}

	if forecast.Condition.Pressure != 1008 {
		t.Error("Forecast should properly return Condition.Pressure")
	}
}

func TestForecastWithErr(t *testing.T) {
	server, c := testTools(200, "{foo")
	defer server.Close()
	_, err := c.Forecast("123")

	if err.Error() != "invalid character 'f' looking for beginning of object key string" {
		t.Error("Forecast should properly catch and return errors")
	}
}

func TestForecastWithNonOKResp(t *testing.T) {
	server, c := testTools(500, resp)
	defer server.Close()
	_, err := c.Forecast("123")

	expected := "http://magicseaweed.com/api/fakeKey/forecast/?spot_id=123 returned HTTP status code 500"
	if err.Error() != expected {
		t.Error(fmt.Sprintf("expected error '%s'; received '%s'", expected, err.Error()))
	}
}

func TestWeekendNoForecast(t *testing.T) {
	server, c := testTools(200, resp)
	defer server.Close()
	forecasts, _ := c.Weekend("123")

	if len(forecasts) > 0 {
		t.Error("Weekend forecasts should be empty")
	}
}
