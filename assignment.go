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
	err := c.request("GET", "/v2/assignments/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

func (c *Client) AssignmentList(params *AssignmentListParams) (*AssignmentPage, error) {
	obj := &AssignmentPage{}
	err := c.request("GET", "/v2/assignments", params, nil, obj)
	return obj, err
}

func (c *Client) AssignmentStart(params *AssignmentStartParams) (*Assignment, error) {
	obj := &Assignment{}
	err := c.request("POST", "/v2/assignments/"+strconv.Itoa(int(*params.ID))+"/start", params, params, obj)
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
	AvailableAt   *time.Time   `json:"available_at"`
	BurnedAt      *time.Time   `json:"burned_at"`
	CreatedAt     time.Time    `json:"created_at"`
	Hidden        bool         `json:"hidden"`
	PassedAt      *time.Time   `json:"passed_at"`
	ResurrectedAt *time.Time   `json:"resurrected_at"`
	SRSStage      int          `json:"srs_stage"`
	StartedAt     *time.Time   `json:"started_at"`
	SubjectID     WKID         `json:"subject_id"`
	SubjectType   WKObjectType `json:"subject_type"`
	UnlockedAt    *time.Time   `json:"unlocked_at"`
}

type AssignmentGetParams struct {
	Params
	ID *WKID
}

type AssignmentListParams struct {
	ListParams
	Params

	AvailableAfter                 *WKTime
	AvailableBefore                *WKTime
	Burned                         *bool
	Hidden                         *bool
	IDs                            []WKID
	ImmediatelyAvailableForLessons *bool
	ImmediatelyAvailableForReview  *bool
	InReview                       *bool
	Levels                         []int
	SRSStages                      []int
	Started                        *bool
	SubjectIDs                     []WKID
	Unlocked                       *bool
	UpdatedAfter                   *WKTime
}

func (p *AssignmentListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.AvailableAfter != nil {
		values.Add("available_after", p.AvailableAfter.Encode())
	}

	if p.AvailableBefore != nil {
		values.Add("available_before", p.AvailableBefore.Encode())
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
		values.Add("updated_after", p.UpdatedAfter.Encode())
	}

	return values.Encode()
}

type AssignmentPage struct {
	*PageObject
	Data []*Assignment `json:"data"`
}

type AssignmentStartParams struct {
	Params
	ID        *WKID   `json:"-"`
	StartedAt *WKTime `json:"started_at"`
}
