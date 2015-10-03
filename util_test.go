package seaweed

import "testing"

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
		Forecast{
			Timestamp:      1, // 9/29
			LocalTimestamp: 1443571200,
		},
		Forecast{
			Timestamp:      2,
			LocalTimestamp: 1443872436, // 10/3
		},
	}

	if len(matchDays(forecasts, 29)) != 1 {
		t.Error("matchDays should properly return only the matched forecast days for a date")
	}

	if matchDays(forecasts, 29)[0].Timestamp != 1 {
		t.Error("matchDays should properly return the matched forecast days for a date")
	}

	if matchDays(forecasts, 3)[0].Timestamp != 2 {
		t.Error("matchDays should properly return the matched forecast days for a date")
	}
}
