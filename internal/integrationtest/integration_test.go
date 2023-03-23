package integrationtest

import (
	"os"
	"testing"
	"time"

	"github.com/mdb/seaweed"
)

var client *seaweed.Client

func TestMain(m *testing.M) {
	client = seaweed.NewClient(os.Getenv("MAGIC_SEAWEED_API_KEY"))
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestForecast_Integration(t *testing.T) {
	resp, err := client.Forecast("391")
	if err != nil {
		t.Error(err)
	}

	if len(resp) == 0 {
		t.Error("Forecast returned no forecasts")
	}

	if resp[0].LocalTimestamp == 0 {
		t.Error("Forecast returned no forecast timestamp")
	}
}

func TestToday_Integration(t *testing.T) {
	resp, err := client.Today("391")
	if err != nil {
		t.Error(err)
	}

	if len(resp) == 0 {
		t.Error("Today returned no forecasts")
	}

	today := time.Now().UTC()

	for _, forecast := range resp {
		fd := time.Unix(forecast.LocalTimestamp, 0).UTC()

		if fd.Day() != today.Day() {
			t.Errorf("Today returned forecast for '%s'", fd.String())
		}
	}
}

func TestTomorrow_Integration(t *testing.T) {
	resp, err := client.Tomorrow("391")
	if err != nil {
		t.Error(err)
	}

	if len(resp) == 0 {
		t.Error("API returned no forecasts")
	}

	tomorrow := time.Now().UTC().AddDate(0, 0, 1)

	for _, forecast := range resp {
		fd := time.Unix(forecast.LocalTimestamp, 0).UTC()

		if fd.Day() != tomorrow.Day() {
			t.Errorf("Tomorrow returned forecast for '%s'", fd.String())
		}
	}
}

func TestWeekend_Integration(t *testing.T) {
	resp, err := client.Weekend("391")
	if err != nil {
		t.Error(err)
	}

	if len(resp) == 0 {
		t.Error("API returned no forecasts")
	}

	for _, forecast := range resp {
		fd := time.Unix(forecast.LocalTimestamp, 0).UTC().Weekday().String()

		if fd != "Saturday" && fd != "Sunday" {
			t.Errorf("Weekend returned forecast for '%s'", fd)
		}
	}
}
