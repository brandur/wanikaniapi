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

// ReviewCreate creates a review for a specific AssignmentID. Using the related
// subject_id is also a valid alternative to using AssignmentID.
//
// Some criteria must be met in order for a review to be created: AvailableAt
// must be not null and in the past.
//
// When a review is registered, the associated assignment and review statistic
// are both updated. These are returned in the response body under
// ResourcesUpdated.
func (c *Client) ReviewCreate(params *ReviewCreateParams) (*Review, error) {
	wrapper := &reviewCreateParamsWrapper{Params: params.Params, Review: params}
	obj := &Review{}
	err := c.request("POST", "/v2/reviews", params, wrapper, obj)
	return obj, err
}

// ReviewGet retrieves a specific review by its ID.
func (c *Client) ReviewGet(params *ReviewGetParams) (*Review, error) {
	obj := &Review{}
	err := c.request("GET", "/v2/reviews/"+strconv.FormatInt(int64(*params.ID), 10), params, nil, obj)
	return obj, err
}

// ReviewList returns a collection of all reviews, ordered by ascending
// CreatedAt, 1000 at a time.
func (c *Client) ReviewList(params *ReviewListParams) (*ReviewPage, error) {
	obj := &ReviewPage{}
	err := c.request("GET", "/v2/reviews", params, nil, obj)
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

// Review logs all the correct and incorrect answers provided through the
// 'Reviews' section of WaniKani. Review records are created when a user
// answers all the parts of a subject correctly once; some subjects have both
// meaning or reading parts, and some only have one or the other. Note that
// reviews are not created for the quizzes in lessons.
type Review struct {
	Object
	Data *ReviewData `json:"data"`
}

// ReviewCreateParams are parameters for ReviewCreate.
type ReviewCreateParams struct {
	Params
	AssignmentID            *WKID   `json:"assignment_id,omitempty"`
	CreatedAt               *WKTime `json:"created_at,omitempty"`
	IncorrectMeaningAnswers *int    `json:"incorrect_meaning_answers,omitempty"`
	IncorrectReadingAnswers *int    `json:"incorrect_reading_answers,omitempty"`
	SubjectID               *WKID   `json:"subject_id,omitempty"`
}

type reviewCreateParamsWrapper struct {
	Params
	Review *ReviewCreateParams `json:"review"`
}

// ReviewData contains core data of Review.
type ReviewData struct {
	AssignmentID             WKID      `json:"assignment_id"`
	CreatedAt                time.Time `json:"created_at"`
	EndingSRSStage           int       `json:"ending_srs_stage"`
	IncorrectMeaningAnswers  int       `json:"incorrect_meaning_answers"`
	IncorrectReadingAnswers  int       `json:"incorrect_reading_answers"`
	SpacedRepetitionSystemID WKID      `json:"spaced_repetition_system_id"`
	StartingSRSStage         int       `json:"starting_srs_stage"`
	SubjectID                WKID      `json:"subject_id"`
}

// ReviewGetParams are parameters for ReviewGet.
type ReviewGetParams struct {
	Params
	ID *WKID
}

// ReviewListParams are parameters for ReviewList.
type ReviewListParams struct {
	ListParams
	Params

	AssignmentIDs []WKID
	IDs           []WKID
	SubjectIDs    []WKID
	UpdatedAfter  *WKTime
}

// EncodeToQuery encodes parametes to a query string.
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
		values.Add("updated_after", p.UpdatedAfter.Encode())
	}

	return values.Encode()
}

// ReviewPage represents a single page of Reviews.
type ReviewPage struct {
	PageObject
	Data []*Review `json:"data"`
}
