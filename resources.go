package seaweed

type Day struct {
	Timestamp      int64 `json:"timestamp,omitempty"`
	LocalTimestamp int64 `json:"localTimestamp,omitempty"`
	IssueTimestamp int64 `json:"issueTimestamp,omitempty"`
	FadedRating    int64 `json:"FadedRating,omitempty"`
	SolidRating    int64 `json:"SolidRating,omitempty"`
}
