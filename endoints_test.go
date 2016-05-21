package seaweed

import "testing"

func TestSpotEp(t *testing.T) {
	if spotEp(c, "123") != "http://magicseaweed.com/api/fakeKey/forecast/?spot_id=123" {
		t.Error("spotEp should return the proper endpoint")
	}
}
