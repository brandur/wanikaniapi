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

func (c *Client) ResetGet(params *ResetGetParams) (*Reset, error) {
	obj := &Reset{}
	err := c.request("GET", "/v2/resets/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

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

type Reset struct {
	Object
	Data *ResetData `json:"data"`
}

type ResetData struct {
	Confirmed_at  *time.Time `json:"confirmed_at"`
	CreatedAt     time.Time  `json:"created_at"`
	OriginalLevel int        `json:"original_evel"`
	TargetLevel   int        `json:"target_level"`
}

type ResetGetParams struct {
	*Params
	ID *WKID
}

type ResetListParams struct {
	*ListParams
	IDs          []WKID
	UpdatedAfter *WKTime
}

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

type ResetPage struct {
	*PageObject
	Data []*Reset `json:"data"`
}
