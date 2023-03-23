package seaweed

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	logging "github.com/op/go-logging"
)

var resp string

func TestMain(m *testing.M) {
	content, err := ioutil.ReadFile("testdata/response.json")
	if err != nil {
		log.Fatal(err)
	}

	resp = string(content)

	exitVal := m.Run()
	os.Exit(exitVal)
}

type testClock struct{}

func (testClock) Now() time.Time {
	return time.Unix(1442355356, 0).UTC()
}

func testServerAndClient(code int, body string) (*httptest.Server, *Client) {
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

	client := &Client{
		"fakeKey",
		httpClient,
		NewLogger(logging.INFO),
		testClock{},
	}

	return server, client
}

func TestNewClient(t *testing.T) {
	client := NewClient("fakeKey")

	if client.APIKey != "fakeKey" {
		t.Error("NewClient should properly set the API key")
	}
}

func TestForecast(t *testing.T) {
	tests := []struct {
		desc                 string
		body                 string
		code                 int
		expectError          error
		expectForecastCount  int
		expectLocalTimestamp int64
	}{{
		desc:                 "when successful",
		body:                 resp,
		code:                 200,
		expectForecastCount:  3,
		expectLocalTimestamp: 1442355356,
	}, {
		desc:                "when the response body is invalid JSON",
		body:                "{foo:",
		code:                200,
		expectForecastCount: 0,
		expectError:         errors.New("invalid character 'f' looking for beginning of object key string"),
	}, {
		desc:                "when the response code is not OK",
		body:                resp,
		code:                500,
		expectForecastCount: 0,
		expectError:         errors.New("http://magicseaweed.com/api/fakeKey/forecast/?spot_id=123 returned HTTP status code 500"),
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			server, c := testServerAndClient(test.code, test.body)
			defer server.Close()
			forecasts, err := c.Forecast("123")

			if err != nil && test.expectError == nil {
				t.Errorf("expected '%s' not to error; got '%v'", test.desc, err)
			}

			if test.expectError != nil && err == nil {
				t.Errorf("expected error '%s'; got '%v'", test.expectError.Error(), err)
			}

			if test.expectError != nil && err != nil && test.expectError.Error() != err.Error() {
				t.Errorf("expected error '%s'; got '%v'", test.expectError.Error(), err)
			}

			if len(forecasts) != test.expectForecastCount {
				t.Errorf("expected '%d' forecasts; got '%d'", test.expectForecastCount, len(forecasts))
			}

			if test.expectError == nil && err == nil {
				if forecasts[0].LocalTimestamp != test.expectLocalTimestamp {
					t.Errorf("expected LocalTimestamp '%d'; got '%d'", test.expectLocalTimestamp, forecasts[0].LocalTimestamp)
				}
			}
		})
	}
}

func TestWeekend(t *testing.T) {
	tests := []struct {
		desc                 string
		body                 string
		code                 int
		expectError          error
		expectForecastCount  int
		expectLocalTimestamp int64
	}{{
		desc:                 "when successful",
		body:                 resp,
		code:                 200,
		expectForecastCount:  1,
		expectLocalTimestamp: 1677973254,
	}, {
		desc:                "when the response body is invalid JSON",
		body:                "{foo:",
		code:                200,
		expectForecastCount: 0,
		expectError:         errors.New("invalid character 'f' looking for beginning of object key string"),
	}, {
		desc:                "when the response code is not OK",
		body:                resp,
		code:                500,
		expectForecastCount: 0,
		expectError:         errors.New("http://magicseaweed.com/api/fakeKey/forecast/?spot_id=123 returned HTTP status code 500"),
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			server, c := testServerAndClient(test.code, test.body)
			defer server.Close()
			forecasts, err := c.Weekend("123")

			if err != nil && test.expectError == nil {
				t.Errorf("expected '%s' not to error; got '%v'", test.desc, err)
			}

			if test.expectError != nil && err == nil {
				t.Errorf("expected error '%s'; got '%v'", test.expectError.Error(), err)
			}

			if test.expectError != nil && err != nil && test.expectError.Error() != err.Error() {
				t.Errorf("expected error '%s'; got '%v'", test.expectError.Error(), err)
			}

			if len(forecasts) != test.expectForecastCount {
				t.Errorf("expected '%d' forecasts; got '%d'", test.expectForecastCount, len(forecasts))
			}

			if test.expectError == nil && err == nil {
				if forecasts[0].LocalTimestamp != test.expectLocalTimestamp {
					t.Errorf("expected LocalTimestamp '%d'; got '%d'", test.expectLocalTimestamp, forecasts[0].LocalTimestamp)
				}
			}
		})
	}
}

func TestToday(t *testing.T) {
	tests := []struct {
		desc                 string
		body                 string
		code                 int
		expectError          error
		expectForecastCount  int
		expectLocalTimestamp int64
	}{{
		desc:                 "when successful",
		body:                 resp,
		code:                 200,
		expectForecastCount:  1,
		expectLocalTimestamp: 1442355356,
	}, {
		desc:                "when the response body is invalid JSON",
		body:                "{foo:",
		code:                200,
		expectForecastCount: 0,
		expectError:         errors.New("invalid character 'f' looking for beginning of object key string"),
	}, {
		desc:                "when the response code is not OK",
		body:                resp,
		code:                500,
		expectForecastCount: 0,
		expectError:         errors.New("http://magicseaweed.com/api/fakeKey/forecast/?spot_id=123 returned HTTP status code 500"),
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			server, c := testServerAndClient(test.code, test.body)
			defer server.Close()
			forecasts, err := c.Today("123")

			if err != nil && test.expectError == nil {
				t.Errorf("expected '%s' not to error; got '%v'", test.desc, err)
			}

			if test.expectError != nil && err == nil {
				t.Errorf("expected error '%s'; got '%v'", test.expectError.Error(), err)
			}

			if test.expectError != nil && err != nil && test.expectError.Error() != err.Error() {
				t.Errorf("expected error '%s'; got '%v'", test.expectError.Error(), err)
			}

			if len(forecasts) != test.expectForecastCount {
				t.Errorf("expected '%d' forecasts; got '%d'", test.expectForecastCount, len(forecasts))
			}

			if test.expectError == nil && err == nil {
				if forecasts[0].LocalTimestamp != test.expectLocalTimestamp {
					t.Errorf("expected LocalTimestamp '%d'; got '%d'", test.expectLocalTimestamp, forecasts[0].LocalTimestamp)
				}
			}
		})
	}
}

func TestTomorrow(t *testing.T) {
	tests := []struct {
		desc                 string
		body                 string
		code                 int
		expectError          error
		expectForecastCount  int
		expectLocalTimestamp int64
	}{{
		desc:                 "when successful",
		body:                 resp,
		code:                 200,
		expectForecastCount:  1,
		expectLocalTimestamp: 1442441756,
	}, {
		desc:                "when the response body is invalid JSON",
		body:                "{foo:",
		code:                200,
		expectForecastCount: 0,
		expectError:         errors.New("invalid character 'f' looking for beginning of object key string"),
	}, {
		desc:                "when the response code is not OK",
		body:                resp,
		code:                500,
		expectForecastCount: 0,
		expectError:         errors.New("http://magicseaweed.com/api/fakeKey/forecast/?spot_id=123 returned HTTP status code 500"),
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			server, c := testServerAndClient(test.code, test.body)
			defer server.Close()
			forecasts, err := c.Tomorrow("123")

			if err != nil && test.expectError == nil {
				t.Errorf("expected '%s' not to error; got '%v'", test.desc, err)
			}

			if test.expectError != nil && err == nil {
				t.Errorf("expected error '%s'; got '%v'", test.expectError.Error(), err)
			}

			if test.expectError != nil && err != nil && test.expectError.Error() != err.Error() {
				t.Errorf("expected error '%s'; got '%v'", test.expectError.Error(), err)
			}

			if len(forecasts) != test.expectForecastCount {
				t.Errorf("expected '%d' forecasts; got '%d'", test.expectForecastCount, len(forecasts))
			}

			if test.expectError == nil && err == nil {
				if forecasts[0].LocalTimestamp != test.expectLocalTimestamp {
					t.Errorf("expected LocalTimestamp '%d'; got '%d'", test.expectLocalTimestamp, forecasts[0].LocalTimestamp)
				}
			}
		})
	}
}
