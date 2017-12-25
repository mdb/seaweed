package seaweed

// Forecast represents a Seaweed API forecast
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

// Swell represents a Seaweed API forecast's swell
type Swell struct {
	MinBreakingHeight    int        `json:"minBreakingHeight"`
	AbsMinBreakingHeight float64    `json:"absMinBreakingHeight"`
	MaxBreakingHeight    int        `json:"maxBreakingHeight"`
	AbsMaxBreakingHeight float64    `json:"absMaxBreakingHeight"`
	Unit                 string     `json:"unit"`
	Components           Components `json:"components"`
}

// Components represents a Seaweed API forecast's swell's components
type Components struct {
	Combined Component `json:"combined"`
	Primary  Component `json:"primary"`
}

// Component represents a Seaweed API forecast's swell component
type Component struct {
	Height           float64 `json:"height"`
	Period           int     `json:"period"`
	Direction        float64 `json:"direction"`
	CompassDirection string  `json:"compassDirection"`
}

// Wind represents a Seaweed API forecast's wind
type Wind struct {
	Speed            int    `json:"speed"`
	Direction        int64  `json:"direction"`
	CompassDirection string `json:"compassDirection"`
	Chill            int64  `json:"chill"`
	Gusts            int64  `json:"gusts"`
	Unit             string `json:"unit"`
}

// Condition represents a Seaweed API forecast's condition
type Condition struct {
	Pressure     int64  `json:"pressure"`
	Temperature  int64  `json:"temperature"`
	Weather      string `json:"weather"`
	Unit         string `json:"f"`
	UnitPressure string `json:"unitPressure"`
}
