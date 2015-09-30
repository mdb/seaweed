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
		t.Error("NewClient should properly set the API key")
	}
}
