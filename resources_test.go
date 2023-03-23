package seaweed

import (
	"testing"
	"time"
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

func TestForecast_IsDay(t *testing.T) {
	ts := int64(1677973254)
	f := Forecast{
		LocalTimestamp: ts,
	}

	today := time.Unix(ts, 0)
	tomorrow := time.Unix(ts, 0).AddDate(0, 0, 1)

	if !f.IsDay(today) {
		t.Error("IsDay should return true if a forecast pertains to the day it's passed")
	}

	if f.IsDay(tomorrow) {
		t.Error("IsDay should return false if a forecast does not pertain to the day it's passed")
	}
}
