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

// ResetGet retrieves a specific reset by its ID.
func (c *Client) ResetGet(params *ResetGetParams) (*Reset, error) {
	obj := &Reset{}
	err := c.request("GET", "/v2/resets/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

// ResetList returns a collection of all resets, ordered by ascending
// CreatedAt, 500 at a time.
func (c *Client) ResetList(params *ResetListParams) (*ResetPage, error) {
	obj := &ResetPage{}
	err := c.request("GET", "/v2/resets", params, nil, obj)
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

// Reset represents a user reset.
//
// Users can reset their progress back to any level at or below their current
// level. When they reset to a particular level, all of the assignments and
// review_statistics at that level or higher are set back to their default
// state.
//
// Resets contain information about when those resets happen, the starting
// level, and the target level.
type Reset struct {
	Object
	Data *ResetData `json:"data"`
}

// ResetData contains core data of Reset.
type ResetData struct {
	ConfirmedAt   *time.Time `json:"confirmed_at"`
	CreatedAt     time.Time  `json:"created_at"`
	OriginalLevel int        `json:"original_evel"`
	TargetLevel   int        `json:"target_level"`
}

// ResetGetParams are parameters for ResetGet.
type ResetGetParams struct {
	Params
	ID *WKID
}

// ResetListParams are parameters for ResetList.
type ResetListParams struct {
	ListParams
	Params

	IDs          []WKID
	UpdatedAfter *WKTime
}

// EncodeToQuery encodes parametes to a query string.
func (p *ResetListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}
	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Encode())
	}

	return values.Encode()
}

// ResetPage represents a single page of Resets.
type ResetPage struct {
	PageObject
	Data []*Reset `json:"data"`
}
