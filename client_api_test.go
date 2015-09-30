package seaweed

import "testing"

var c = NewClient("fakeKey")

func TestTomorrow(t *testing.T) {
	tomorrow, _ := c.Tomorrow()

	if tomorrow.Date != "foo" {
		t.Error("NewClient should properly set the API key")
	}
}
