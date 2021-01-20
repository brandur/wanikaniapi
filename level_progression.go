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

// LevelProgressionGet retrieves a specific level progression by its ID.
func (c *Client) LevelProgressionGet(params *LevelProgressionGetParams) (*LevelProgression, error) {
	obj := &LevelProgression{}
	err := c.request("GET", "/v2/level_progressions/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

// LevelProgressionList returns a collection of all level progressions, ordered
// by ascending CreatedAt, 500 at a time.
func (c *Client) LevelProgressionList(params *LevelProgressionListParams) (*LevelProgressionPage, error) {
	obj := &LevelProgressionPage{}
	err := c.request("GET", "/v2/level_progressions", params, nil, obj)
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

// LevelProgression contains information about a user's progress through the
// WaniKani levels.
type LevelProgression struct {
	Object
	Data *LevelProgressionData `json:"data"`
}

// LevelProgressionData contains core data of LevelProgression.
type LevelProgressionData struct {
	AbandonedAt *time.Time `json:"abandoned_at"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time  `json:"created_at"`
	Level       int        `json:"level"`
	PassedAt    *time.Time `json:"passed_at"`
	StartedAt   *time.Time `json:"started_at"`
	UnlockedAt  *time.Time `json:"unlocked_at"`
}

// LevelProgressionGetParams are parameters for LevelProgressionGet.
type LevelProgressionGetParams struct {
	Params
	ID *WKID
}

// LevelProgressionListParams are parameters for LevelProgressionList.
type LevelProgressionListParams struct {
	ListParams
	Params

	IDs          []WKID
	UpdatedAfter *WKTime
}

// EncodeToQuery encodes parametes to a query string.
func (p *LevelProgressionListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}

	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Encode())
	}

	return values.Encode()
}

// LevelProgressionPage represents a single page of LevelProgressions.
type LevelProgressionPage struct {
	PageObject
	Data []*LevelProgression `json:"data"`
}
