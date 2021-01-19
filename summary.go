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

type Summary struct {
	Object
	Data *SummaryData `json:"data"`
}

type SummaryData struct {
	Lessons       []*SummaryLesson `json:"lessons"`
	NextReviewsAt *time.Time       `json:"next_reviews_at"`
	Reviews       []*SummaryReview `json:"reviews"`
}

type SummaryGetParams struct {
	Params
}

type SummaryLesson struct {
	AvailableAt time.Time `json:"available_at"`
	SubjectIDs  []WKID    `json:"subject_ids"`
}

type SummaryPage struct {
	*PageObject
	Data []*Summary `json:"data"`
}

type SummaryReview struct {
	AvailableAt time.Time `json:"available_at"`
	SubjectIDs  []WKID    `json:"subject_ids"`
}
