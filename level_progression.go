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

func (c *Client) LevelProgressionGet(params *LevelProgressionGetParams) (*LevelProgression, error) {
	obj := &LevelProgression{}
	err := c.request("GET", "/v2/level_progressions/"+strconv.Itoa(int(*params.ID)), "", nil, obj)
	return obj, err
}

func (c *Client) LevelProgressionList(params *LevelProgressionListParams) (*LevelProgressionPage, error) {
	obj := &LevelProgressionPage{}
	err := c.request("GET", "/v2/level_progressions", params.EncodeToQuery(), nil, obj)
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

type LevelProgression struct {
	Object
	Data *LevelProgressionData `json:"data"`
}

type LevelProgressionData struct {
	CreatedAt   time.Time `json:"created_at"`
	Level       int       `json:"level"`
	UnlockedAt  time.Time `json:"unlocked_at"`
	StartedAt   time.Time `json:"started_at"`
	PassedAt    time.Time `json:"passed_at"`
	CompletedAt time.Time `json:"completed_at"`
	AbandonedAt time.Time `json:"abandoned_at"`
}

type LevelProgressionGetParams struct {
	ID *ID
}

type LevelProgressionListParams struct {
	*ListParams
	IDs          []ID
	UpdatedAfter *time.Time
}

func (p *LevelProgressionListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}

	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Format(time.RFC3339))
	}

	return values.Encode()
}

type LevelProgressionPage struct {
	*PageObject
	Data []*LevelProgression `json:"data"`
}
