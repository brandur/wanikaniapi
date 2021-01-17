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

func (c *Client) AssignmentGet(params *AssignmentGetParams) (*Assignment, error) {
	obj := &Assignment{}
	err := c.request("GET", "/v2/assignments/"+strconv.Itoa(int(*params.ID)), "", nil, obj)
	return obj, err
}

func (c *Client) AssignmentList(params *AssignmentListParams) (*AssignmentPage, error) {
	obj := &AssignmentPage{}
	err := c.request("GET", "/v2/assignments", params.EncodeToQuery(), nil, obj)
	return obj, err
}

func (c *Client) AssignmentStart(params *AssignmentStartParams) (*Assignment, error) {
	obj := &Assignment{}
	err := c.request("POST", "/v2/assignments/"+strconv.Itoa(int(*params.ID))+"/start", "", params, obj)
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

type Assignment struct {
	Object
	Data *AssignmentData `json:"data"`
}

type AssignmentData struct {
	AvailableAt   *time.Time `json:"available_at"`
	BurnedAt      *time.Time `json:"burned_at"`
	CreatedAt     time.Time  `json:"created_at"`
	Hidden        bool       `json:"hidden"`
	PassedAt      *time.Time `json:"passed_at"`
	ResurrectedAt *time.Time `json:"resurrected_at"`
	SRSStage      int        `json:"srs_stage"`
	StartedAt     *time.Time `json:"started_at"`
	SubjectID     ID         `json:"subject_id"`
	SubjectType   ObjectType `json:"subject_type"`
	UnlockedAt    *time.Time `json:"unlocked_at"`
}

type AssignmentGetParams struct {
	ID *ID
}

type AssignmentListParams struct {
	*ListParams
	AvailableAfter                 *time.Time
	AvailableBefore                *time.Time
	Burned                         *bool
	Hidden                         *bool
	IDs                            []ID
	ImmediatelyAvailableForLessons *bool
	ImmediatelyAvailableForReview  *bool
	InReview                       *bool
	Levels                         []int
	SRSStages                      []int
	Started                        *bool
	SubjectIDs                     []ID
	Unlocked                       *bool
	UpdatedAfter                   *time.Time
}

func (p *AssignmentListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.AvailableAfter != nil {
		values.Add("available_after", p.AvailableAfter.Format(time.RFC3339))
	}

	if p.AvailableBefore != nil {
		values.Add("available_before", p.AvailableBefore.Format(time.RFC3339))
	}

	if p.Burned != nil {
		values.Add("burned", strconv.FormatBool(*p.Burned))
	}

	if p.Hidden != nil {
		values.Add("hidden", strconv.FormatBool(*p.Hidden))
	}

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}

	if p.ImmediatelyAvailableForLessons != nil {
		values.Add("immediately_available_for_lessons",
			strconv.FormatBool(*p.ImmediatelyAvailableForLessons))
	}

	if p.ImmediatelyAvailableForReview != nil {
		values.Add("immediately_available_for_review",
			strconv.FormatBool(*p.ImmediatelyAvailableForReview))
	}

	if p.InReview != nil {
		values.Add("in_review", strconv.FormatBool(*p.InReview))
	}

	if p.Levels != nil {
		values.Add("levels", joinInts(p.Levels, ","))
	}

	if p.SRSStages != nil {
		values.Add("srs_stages", joinInts(p.SRSStages, ","))
	}

	if p.Started != nil {
		values.Add("started", strconv.FormatBool(*p.Started))
	}

	if p.Unlocked != nil {
		values.Add("unlocked", strconv.FormatBool(*p.Unlocked))
	}

	if p.SubjectIDs != nil {
		values.Add("subject_ids", joinIDs(p.SubjectIDs, ","))
	}

	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Format(time.RFC3339))
	}

	return values.Encode()
}

type AssignmentPage struct {
	*PageObject
	Data []*Assignment `json:"data"`
}

type AssignmentStartParams struct {
	ID        *ID     `json:"-"`
	StartedAt *WKTime `json:"started_at"`
}
