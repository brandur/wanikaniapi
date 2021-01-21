package wanikaniapi

import (
	"time"
)

//////////////////////////////////////////////////////////////////////////////
//
//
//
// Exported functions
//
//
//
//////////////////////////////////////////////////////////////////////////////

// SummaryGet retrieves a summary report.
func (c *Client) SummaryGet(params *SummaryGetParams) (*Summary, error) {
	obj := &Summary{}
	err := c.request("GET", "/v2/summary", params, nil, obj)
	return obj, err
}

//////////////////////////////////////////////////////////////////////////////
//
//
//
// Exported constants/types
//
//
//
//////////////////////////////////////////////////////////////////////////////

// Summary contains currently available lessons and reviews and the reviews
// that will become available in the next 24 hours, grouped by the hour.
type Summary struct {
	Object
	Data *SummaryData `json:"data"`
}

// SummaryData contains core data of Summary.
type SummaryData struct {
	Lessons       []*SummaryLesson `json:"lessons"`
	NextReviewsAt *time.Time       `json:"next_reviews_at"`
	Reviews       []*SummaryReview `json:"reviews"`
}

// SummaryGetParams are parameters for SummaryGet.
type SummaryGetParams struct {
	Params
}

// SummaryLesson provides a summary about lessons.
type SummaryLesson struct {
	AvailableAt time.Time `json:"available_at"`
	SubjectIDs  []WKID    `json:"subject_ids"`
}

// SummaryPage represents a single page of Summaries.
type SummaryPage struct {
	PageObject
	Data []*Summary `json:"data"`
}

// SummaryReview provides a summary about reviews.
type SummaryReview struct {
	AvailableAt time.Time `json:"available_at"`
	SubjectIDs  []WKID    `json:"subject_ids"`
}
