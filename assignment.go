package wanikaniapi

import (
	"strconv"
	"time"
)

const (
	ObjectTypeAssignment = ObjectType("assigment")
)

type Assignment struct {
	Object
	Data *AssignmentData `json:"data"`
}

type AssignmentData struct {
	AvailableAt   *time.Time `json:"available_at"`
	BurnedAt      *time.Time `json:"burned_at"`
	CreatedAt     time.Time  `json:"created_at"`
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
	Levels []int
}

func (p *AssignmentListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.Levels != nil {
		values.Add("levels", joinInts(p.Levels, ","))
	}

	return values.Encode()
}

type AssignmentPage struct {
	*PageObject
	Data []*Assignment `json:"data"`
}

func (c *Client) GetAssignment(params *AssignmentGetParams) (*Assignment, error) {
	obj := &Assignment{}
	err := c.request("GET", "/v2/assignments/"+strconv.Itoa(int(*params.ID)), "", obj)
	return obj, err
}

func (c *Client) ListAssignments(params *AssignmentListParams) (*AssignmentPage, error) {
	obj := &AssignmentPage{}
	err := c.request("GET", "/v2/assignments", params.EncodeToQuery(), obj)
	return obj, err
}
