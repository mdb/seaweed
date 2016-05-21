package seaweed

import (
	"os"
	"testing"
	"time"
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
