package integrationtest

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/mdb/seaweed"
)

const envVarName string = "MAGIC_SEAWEED_API_KEY"

var client *seaweed.Client

func TestMain(m *testing.M) {
	key := os.Getenv(envVarName)
	if key == "" {
		log.Fatalf("%s environment variable not set", envVarName)
	}

	client = seaweed.NewClient(key)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGet_Integration(t *testing.T) {
	resp, err := seaweed.Get(os.Getenv(envVarName), "391")
	if err != nil {
		t.Error(err)
	}

	if len(resp) == 0 {
		t.Error("Get returned no forecasts")
	}

	if resp[0].LocalTimestamp == 0 {
		t.Error("Get returned no forecast timestamp")
	}
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

func TestForecast_Integration_error(t *testing.T) {
	c := seaweed.NewClient("")
	resp, err := c.Forecast("391")
	expected := "Unable to authenticate request: Ensure your API key is passed correctly. Refer to the API docs."
	if err.Error() != expected {
		t.Errorf("expected Forecast to err with '%s'; got '%s'", expected, err.Error())
	}

	if len(resp) > 0 {
		t.Error("erroring Forecast returned forecasts")
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

	for _, forecast := range resp {
		fd := time.Unix(forecast.LocalTimestamp, 0).UTC().Weekday().String()

		if fd != "Saturday" && fd != "Sunday" {
			t.Errorf("Weekend returned forecast for '%s'", fd)
		}
	}
}
