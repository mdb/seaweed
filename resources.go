package seaweed

type Day struct {
	Timestamp      int64 `json:"timestamp"`
	LocalTimestamp int64 `json:"localTimestamp"`
	IssueTimestamp int64 `json:"issueTimestamp"`
	FadedRating    int   `json:"FadedRating"`
	SolidRating    int   `json:"SolidRating"`
	Swell          Swell `json:"swell"`
}

type Swell struct {
	MinBreakingHeight    int     `json:"minBreakingHeight"`
	AbsMinBreakingHeight float64 `json:"absMinBreakingHeight"`
	MaxBreakingHeight    int     `json:"maxBreakingHeight"`
	AbsMaxBreakingHeight float64 `json:"absMaxBreakingHeight"`
	Unit                 string  `json:"unit"`
}
