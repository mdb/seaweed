package seaweed

type Day struct {
	Timestamp      int64 `json:"timestamp"`
	LocalTimestamp int64 `json:"localTimestamp"`
	IssueTimestamp int64 `json:"issueTimestamp"`
	FadedRating    int   `json:"FadedRating"`
	SolidRating    int   `json:"SolidRating"`
	Swell          Swell `json:"swell"`
	Wind           Wind  `json:"wind"`
}

type Swell struct {
	MinBreakingHeight    int        `json:"minBreakingHeight"`
	AbsMinBreakingHeight float64    `json:"absMinBreakingHeight"`
	MaxBreakingHeight    int        `json:"maxBreakingHeight"`
	AbsMaxBreakingHeight float64    `json:"absMaxBreakingHeight"`
	Unit                 string     `json:"unit"`
	Components           Components `json:"components"`
}

type Components struct {
	Combined Component `json:"combined"`
	Primary  Component `json:"primary"`
}

type Component struct {
	Height           float64 `json:"height"`
	Period           int     `json:"period"`
	Direction        float64 `json:"direction"`
	CompassDirection string  `json:"compassDirection"`
}

type Wind struct {
	Speed            int    `json:"speed"`
	Direction        int64  `json:"direction"`
	CompassDirection string `json:"compassDirection"`
	Chill            int64  `json:"chill"`
	Gusts            int64  `json:"gusts"`
	Unit             string `json:"mph"`
}
