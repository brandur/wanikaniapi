package wanikaniapi

import (
	"strconv"
	"time"
)

type Review struct {
	Object
	Data *ReviewData `json:"data"`
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

func (c *Client) ReviewGet(params *ReviewGetParams) (*Review, error) {
	obj := &Review{}
	err := c.request("GET", "/v2/reviews/"+strconv.FormatInt(int64(*params.ID), 10), "", obj)
	return obj, err
}

func (c *Client) ReviewList(params *ReviewListParams) (*ReviewPage, error) {
	obj := &ReviewPage{}
	err := c.request("GET", "/v2/reviews", params.EncodeToQuery(), obj)
	return obj, err
}
