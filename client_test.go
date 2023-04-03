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
	"reflect"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	resp      string
	errorResp string
)

func TestMain(m *testing.M) {
	content, err := ioutil.ReadFile("testdata/response.json")
	if err != nil {
		log.Fatal(err)
	}

	resp = string(content)

	errContent, err := ioutil.ReadFile("testdata/error.json")
	if err != nil {
		log.Fatal(err)
	}

	errorResp = string(errContent)

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

	client := NewClient(
		"fakeKey",
		WithBaseURL(server.URL),
		WithHTTPClient(httpClient),
		WithClock(testClock{}),
	)

	return server, client
}

func TestNewClient(t *testing.T) {
	client := NewClient("fakeKey")

	if reflect.TypeOf(client.Logger) != reflect.TypeOf(&logrus.Logger{}) {
		t.Error("NewClient should properly set the Logger")
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
		expectError:         errors.New("unexpected API response '{foo:': invalid character 'f' looking for beginning of object key string"),
	}, {
		desc:                "when the response code is not OK",
		body:                resp,
		code:                500,
		expectForecastCount: 0,
		expectError:         errors.New("GET /api/<REDACTED>/forecast/?spot_id=123 returned HTTP status code 500"),
	}, {
		desc:                "when the response code is OK but the response body specifies an error",
		body:                errorResp,
		code:                200,
		expectForecastCount: 0,
		expectError:         errors.New("Unable to authenticate request: Ensure your API key is passed correctly. Refer to the API docs."),
	}, {
		desc:                "when the response code is OK and the response body indicates an error, but with unexpected JSON",
		body:                "error_response{",
		code:                200,
		expectForecastCount: 0,
		expectError:         errors.New("unexpected API response 'error_response{': invalid character 'e' looking for beginning of value"),
	}}

	for i := range tests {
		test := tests[i]

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

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
		expectError:         errors.New("unexpected API response '{foo:': invalid character 'f' looking for beginning of object key string"),
	}, {
		desc:                "when the response code is not OK",
		body:                resp,
		code:                500,
		expectForecastCount: 0,
		expectError:         errors.New("GET /api/<REDACTED>/forecast/?spot_id=123 returned HTTP status code 500"),
	}}

	for i := range tests {
		test := tests[i]

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

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
		expectError:         errors.New("unexpected API response '{foo:': invalid character 'f' looking for beginning of object key string"),
	}, {
		desc:                "when the response code is not OK",
		body:                resp,
		code:                500,
		expectForecastCount: 0,
		expectError:         errors.New("GET /api/<REDACTED>/forecast/?spot_id=123 returned HTTP status code 500"),
	}}

	for i := range tests {
		test := tests[i]

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

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
		expectError:         errors.New("unexpected API response '{foo:': invalid character 'f' looking for beginning of object key string"),
	}, {
		desc:                "when the response code is not OK",
		body:                resp,
		code:                500,
		expectForecastCount: 0,
		expectError:         errors.New("GET /api/<REDACTED>/forecast/?spot_id=123 returned HTTP status code 500"),
	}}

	for i := range tests {
		test := tests[i]

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

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
