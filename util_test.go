package seaweed

import (
	"testing"
	"time"
)

func TestMatchDays(t *testing.T) {
	forecasts := []Forecast{
		{
			Timestamp:      1, // 9/30
			LocalTimestamp: 1443571200,
		},
		{
			Timestamp:      2,
			LocalTimestamp: 1443872436, // 10/3
		},
	}

	dayOne := time.Unix(forecasts[0].LocalTimestamp, 0).UTC().Day()
	dayTwo := time.Unix(forecasts[1].LocalTimestamp, 0).UTC().Day()

	if len(matchDays(forecasts, dayOne)) != 1 {
		t.Error("matchDays should properly return only the matched forecast days for a date (30)")
	}

	if matchDays(forecasts, dayOne)[0].Timestamp != 1 {
		t.Error("matchDays should properly return the matched forecast days for a date (30)")
	}

	if matchDays(forecasts, dayTwo)[0].Timestamp != 2 {
		t.Error("matchDays should properly return the matched forecast days for a date (3)")
	}
}
