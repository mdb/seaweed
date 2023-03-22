package seaweed

import (
	"testing"
)

func TestForecast_IsWeekend(t *testing.T) {
	saturdayF := Forecast{
		LocalTimestamp: 1677973254,
	}

	if !saturdayF.IsWeekend() {
		t.Error("IsWeekend should properly return return true if a forecast pertains to a Saturday")
	}

	sundayF := Forecast{
		LocalTimestamp: 1678059654,
	}

	if !sundayF.IsWeekend() {
		t.Error("IsWeekend should properly return return true if a forecast pertains to a Sunday")
	}

	weekdayF := Forecast{
		LocalTimestamp: 1677886854,
	}

	if weekdayF.IsWeekend() {
		t.Error("IsWeekend should properly return return false if a forecast pertains to a weekday")
	}
}
