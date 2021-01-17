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

func (c *Client) ReviewCreate(params *ReviewCreateParams) (*Review, error) {
	wrapper := &reviewCreateParamsWrapper{Params: params}
	obj := &Review{}
	err := c.request("POST", "/v2/reviews", "", wrapper, obj)
	return obj, err
}

func (c *Client) ReviewGet(params *ReviewGetParams) (*Review, error) {
	obj := &Review{}
	err := c.request("GET", "/v2/reviews/"+strconv.FormatInt(int64(*params.ID), 10), "", nil, obj)
	return obj, err
}

func (c *Client) ReviewList(params *ReviewListParams) (*ReviewPage, error) {
	obj := &ReviewPage{}
	err := c.request("GET", "/v2/reviews", params.EncodeToQuery(), nil, obj)
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

type Review struct {
	Object
	Data *ReviewData `json:"data"`
}

type ReviewCreateParams struct {
	AssignmentID            *ID     `json:"assignment_id,omitempty"`
	CreatedAt               *WKTime `json:"created_at,omitempty"`
	IncorrectMeaningAnswers *int    `json:"incorrect_meaning_answers,omitempty"`
	IncorrectReadingAnswers *int    `json:"incorrect_reading_answers,omitempty"`
	SubjectID               *ID     `json:"subject_id,omitempty"`
}

type reviewCreateParamsWrapper struct {
	Params *ReviewCreateParams `json:"review"`
}

type ReviewData struct {
	AssignmentID             ID        `json:"assignment_id"`
	CreatedAt                time.Time `json:"created_at"`
	EndingSRSStage           int       `json:"ending_srs_stage"`
	IncorrectMeaningAnswers  int       `json:"incorrect_meaning_answers"`
	IncorrectReadingAnswers  int       `json:"incorrect_reading_answers"`
	SpacedRepetitionSystemID ID        `json:"spaced_repetition_system_id"`
	StartingSRSStage         int       `json:"starting_srs_stage"`
	SubjectID                ID        `json:"subject_id"`
}

type ReviewGetParams struct {
	ID *ID
}

type ReviewListParams struct {
	*ListParams
	AssignmentIDs []ID
	IDs           []ID
	SubjectIDs    []ID
	UpdatedAfter  *time.Time
}

func (p *ReviewListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.AssignmentIDs != nil {
		values.Add("assignment_ids", joinIDs(p.AssignmentIDs, ","))
	}

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}

	if p.SubjectIDs != nil {
		values.Add("subject_ids", joinIDs(p.SubjectIDs, ","))
	}

	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Format(time.RFC3339))
	}

	return values.Encode()
}

type ReviewPage struct {
	*PageObject
	Data []*Review `json:"data"`
}
