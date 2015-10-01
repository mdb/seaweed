package seaweed

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var c = NewClient("fakeKey")

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

	client := &Client{"fakeKey", httpClient}

	return server, client
}

func TestTomorrow(t *testing.T) {
	server, c := testTools(200, resp)
	defer server.Close()
	tomorrow, _ := c.Tomorrow("123")

	if tomorrow.Timestamp != 1443592800 {
		t.Error("Tomorrow should properly return a Timestamp")
	}

	if tomorrow.LocalTimestamp != 1443571200 {
		t.Error("Tomorrow should properly return a LocalTimestamp")
	}

	if tomorrow.IssueTimestamp != 1443592800 {
		t.Error("Tomorrow should properly return an IssueTimestamp")
	}

	if tomorrow.FadedRating != 3 {
		t.Error("Tomorrow should properly return a FadedRating")
	}

	if tomorrow.SolidRating != 0 {
		t.Error("Tomorrow should properly return SolidRating")
	}

	if tomorrow.Swell.MinBreakingHeight != 5 {
		t.Error("Tomorrow should properly return Swell.MinBreakingHeight")
	}

	if tomorrow.Swell.AbsMinBreakingHeight != 4.88 {
		t.Error("Tomorrow should properly return Swell.AbsMinBreakingHeight")
	}

	if tomorrow.Swell.Unit != "ft" {
		t.Error("Tomorrow should properly return Swell.Unit")
	}

	if tomorrow.Swell.MaxBreakingHeight != 8 {
		t.Error("Tomorrow should properly return Swell.MaxBreakingHeight")
	}

	if tomorrow.Swell.AbsMaxBreakingHeight != 7.63 {
		t.Error("Tomorrow should properly return Swell.AbsMaxBreakingHeight")
	}

	if tomorrow.Swell.Components.Combined.Height != 7.5 {
		t.Error("Tomorrow should properly return Swell.Components.Combined.Height")
	}

	if tomorrow.Swell.Components.Primary.Height != 7.5 {
		t.Error("Tomorrow should properly return Swell.Components.Primary.Height")
	}

	if tomorrow.Wind.Speed != 13 {
		t.Error("Tomorrow should properly return Wind.Speed")
	}
}
