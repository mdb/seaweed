package seaweed

import "time"

// APIError represents a Seaweed API error response body.
//
// Note that the Magic Seaweed API may respond with an HTTP status code of 200
// and a response body reporting an error.
type APIError struct {
	ErrorResponse ErrorResponse `json:"error_response"`
}

// ErrorResponse represents a Seaweed API error response.
type ErrorResponse struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"error_msg"`
}

// Forecast represents a Seaweed API forecast.
type Forecast struct {
	Timestamp      int64     `json:"timestamp"`
	LocalTimestamp int64     `json:"localTimestamp"`
	IssueTimestamp int64     `json:"issueTimestamp"`
	FadedRating    int       `json:"FadedRating"`
	SolidRating    int       `json:"SolidRating"`
	Swell          Swell     `json:"swell"`
	Wind           Wind      `json:"wind"`
	Condition      Condition `json:"condition"`
}

// IsWeekend returns true if a forecast pertains to a Saturday or a Sunday.
func (f Forecast) IsWeekend() bool {
	day := time.Unix(f.LocalTimestamp, 0).UTC().Weekday().String()

	if day == "Saturday" || day == "Sunday" {
		return true
	}

	return false
}

// IsDay returns true if a forecast pertains to the day it's passed.
func (f Forecast) IsDay(t time.Time) bool {
	day := time.Unix(f.LocalTimestamp, 0).UTC()

	return day.Day() == t.Day() && day.Month() == t.Month() && day.Year() == t.Year()
}

// Swell represents a Seaweed API forecast's swell.
type Swell struct {
	MinBreakingHeight    int        `json:"minBreakingHeight"`
	AbsMinBreakingHeight float64    `json:"absMinBreakingHeight"`
	MaxBreakingHeight    int        `json:"maxBreakingHeight"`
	AbsMaxBreakingHeight float64    `json:"absMaxBreakingHeight"`
	Unit                 string     `json:"unit"`
	Components           Components `json:"components"`
}

// Components represents a Seaweed API forecast's swell's components.
type Components struct {
	Combined Component `json:"combined"`
	Primary  Component `json:"primary"`
}

// Component represents a Seaweed API forecast's swell component.
type Component struct {
	Height           float64 `json:"height"`
	Period           int     `json:"period"`
	Direction        float64 `json:"direction"`
	CompassDirection string  `json:"compassDirection"`
}

// Wind represents a Seaweed API forecast's wind.
type Wind struct {
	Speed            int    `json:"speed"`
	Direction        int64  `json:"direction"`
	CompassDirection string `json:"compassDirection"`
	Chill            int64  `json:"chill"`
	Gusts            int64  `json:"gusts"`
	Unit             string `json:"unit"`
}

// Condition represents a Seaweed API forecast's condition.
type Condition struct {
	Pressure     int64  `json:"pressure"`
	Temperature  int64  `json:"temperature"`
	Weather      string `json:"weather"`
	Unit         string `json:"unit"`
	UnitPressure string `json:"unitPressure"`
}
