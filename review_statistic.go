package wanikaniapi

import (
	"strconv"
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

// ReviewStatisticGet retrieves a specific review statistic by its ID.
func (c *Client) ReviewStatisticGet(params *ReviewStatisticGetParams) (*ReviewStatistic, error) {
	obj := &ReviewStatistic{}
	err := c.request("GET", "/v2/review_statistics/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

// ReviewStatisticList returns a collection of all review statistics, ordered
// by ascending CreatedAt, 500 at a time.
func (c *Client) ReviewStatisticList(params *ReviewStatisticListParams) (*ReviewStatisticPage, error) {
	obj := &ReviewStatisticPage{}
	err := c.request("GET", "/v2/review_statistics", params, nil, obj)
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

// ReviewStatistic summarizes the activity recorded in reviews. They contain
// sum the number of correct and incorrect answers for both meaning and
// reading. They track current and maximum streaks of correct answers. They
// store the overall percentage of correct answers versus total answers.
//
// A review statistic is created when the user has done their first review on
// the related subject.
type ReviewStatistic struct {
	Object
	Data *ReviewStatisticData `json:"data"`
}

// ReviewStatisticData contains core data of ReviewStatistic.
type ReviewStatisticData struct {
	CreatedAt            time.Time    `json:"created_at"`
	Hidden               bool         `json:"hidden"`
	MeaningCorrect       int          `json:"meaning_correct"`
	MeaningCurrentStreak int          `json:"meaning_current_streak"`
	MeaningIncorrect     int          `json:"meaning_incorrect"`
	MeaningMaxStreak     int          `json:"meaning_max_streak"`
	PercentageCorrect    int          `json:"percentage_correct"`
	ReadingCorrect       int          `json:"reading_correct"`
	ReadingCurrentStreak int          `json:"reading_current_streak"`
	ReadingIncorrect     int          `json:"reading_incorrect"`
	ReadingMaxStreak     int          `json:"reading_max_streak"`
	SubjectID            WKID         `json:"subject_id"`
	SubjectType          WKObjectType `json:"subject_type"`
}

// ReviewStatisticGetParams are parameters for ReviewStatisticGet.
type ReviewStatisticGetParams struct {
	Params
	ID *WKID
}

// ReviewStatisticListParams are parameters for ReviewStatisticList.
type ReviewStatisticListParams struct {
	ListParams
	Params

	Hidden                 *bool
	IDs                    []WKID
	PercentagesGreaterThan *int
	PercentagesLesserThan  *int
	SubjectIDs             []WKID
	SubjectTypes           []WKObjectType
	UpdatedAfter           *WKTime
}

// EncodeToQuery encodes parametes to a query string.
func (p *ReviewStatisticListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.Hidden != nil {
		values.Add("hidden", strconv.FormatBool(*p.Hidden))
	}

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}

	if p.PercentagesGreaterThan != nil {
		values.Add("percentages_greater_than", strconv.Itoa(*p.PercentagesGreaterThan))
	}

	if p.PercentagesLesserThan != nil {
		values.Add("percentages_lesser_than", strconv.Itoa(*p.PercentagesLesserThan))
	}

	if p.SubjectIDs != nil {
		values.Add("subject_ids", joinIDs(p.SubjectIDs, ","))
	}

	if p.SubjectTypes != nil {
		values.Add("subject_types", joinObjectTypes(p.SubjectTypes, ","))
	}

	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Encode())
	}

	return values.Encode()
}

// ReviewStatisticPage represents a single page of ReviewStatistics.
type ReviewStatisticPage struct {
	PageObject
	Data []*ReviewStatistic `json:"data"`
}
