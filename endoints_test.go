package seaweed

import "testing"

func TestSpotEP(t *testing.T) {
	if spotEP(c, "123") != "http://magicseaweed.com/api/fakeKey/forecast/?spot_id=123" {
		t.Error("spotEp should return the proper endpoint")
	}
}
