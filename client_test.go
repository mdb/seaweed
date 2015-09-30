package seaweed

import "testing"

func TestNewClient(t *testing.T) {
	client := NewClient("fakeKey")

	if client.ApiKey != "fakeKey" {
		t.Error("NewClient should properly set the API key")
	}
}
