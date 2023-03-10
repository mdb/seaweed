package seaweed

import (
	"testing"
	"time"
)

func TestSpotEP(t *testing.T) {
	if spotEP(c, "123") != "http://magicseaweed.com/api/fakeKey/forecast/?spot_id=123" {
		t.Error("spotEp should return the proper endpoint")
	}
}

func TestConcat(t *testing.T) {
	joined := concat([]string{
		"foo",
		"bar",
	})

	if joined != "foobar" {
		t.Error("concat should properly concatenate strings")
	}
}

func TestMatchDays(t *testing.T) {
	forecasts := []Forecast{
		{
			Timestamp:      1, // 9/29
			LocalTimestamp: 1443571200,
		},
		{
			Timestamp:      2,
			LocalTimestamp: 1443872436, // 10/3
		},
	}

	dayOne := time.Unix(forecasts[0].LocalTimestamp, 0).Day()
	dayTwo := time.Unix(forecasts[1].LocalTimestamp, 0).Day()

	if len(matchDays(forecasts, dayOne)) != 1 {
		t.Error("matchDays should properly return only the matched forecast days for a date (29)")
	}

	if matchDays(forecasts, dayOne)[0].Timestamp != 1 {
		t.Error("matchDays should properly return the matched forecast days for a date (29)")
	}

	if matchDays(forecasts, dayTwo)[0].Timestamp != 2 {
		t.Error("matchDays should properly return the matched forecast days for a date (3)")
	}
}
