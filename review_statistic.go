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

func (c *Client) ReviewStatisticGet(params *ReviewStatisticGetParams) (*ReviewStatistic, error) {
	obj := &ReviewStatistic{}
	err := c.request("GET", "/v2/review_statistics/"+strconv.Itoa(int(*params.ID)), "", nil, obj)
	return obj, err
}

func (c *Client) ReviewStatisticList(params *ReviewStatisticListParams) (*ReviewStatisticPage, error) {
	obj := &ReviewStatisticPage{}
	err := c.request("GET", "/v2/review_statistics", params.EncodeToQuery(), nil, obj)
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

type ReviewStatistic struct {
	Object
	Data *ReviewStatisticData `json:"data"`
}

type ReviewStatisticData struct {
	CreatedAt            time.Time  `json:"created_at"`
	Hidden               bool       `json:"hidden"`
	MeaningCorrect       int        `json:"meaning_correct"`
	MeaningCurrentStreak int        `json:"meaning_current_streak"`
	MeaningIncorrect     int        `json:"meaning_incorrect"`
	MeaningMaxStreak     int        `json:"meaning_max_streak"`
	PercentageCorrect    int        `json:"percentage_correct"`
	ReadingCorrect       int        `json:"reading_correct"`
	ReadingCurrentStreak int        `json:"reading_current_streak"`
	ReadingIncorrect     int        `json:"reading_incorrect"`
	ReadingMaxStreak     int        `json:"reading_max_streak"`
	SubjectID            ID         `json:"subject_id"`
	SubjectType          ObjectType `json:"subject_type"`
}

type ReviewStatisticGetParams struct {
	ID *ID
}

type ReviewStatisticListParams struct {
	*ListParams
	Hidden                 *bool
	IDs                    []ID
	PercentagesGreaterThan *int
	PercentagesLesserThan  *int
	SubjectIDs             []ID
	SubjectTypes           []ObjectType
	UpdatedAfter           *time.Time
}

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
		values.Add("updated_after", p.UpdatedAfter.Format(time.RFC3339))
	}

	return values.Encode()
}

type ReviewStatisticPage struct {
	*PageObject
	Data []*ReviewStatistic `json:"data"`
}
