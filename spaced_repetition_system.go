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

func (c *Client) SpacedRepetitionSystemGet(params *SpacedRepetitionSystemGetParams) (*SpacedRepetitionSystem, error) {
	obj := &SpacedRepetitionSystem{}
	err := c.request("GET", "/v2/spaced_repetition_systems/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

func (c *Client) SpacedRepetitionSystemList(params *SpacedRepetitionSystemListParams) (*SpacedRepetitionSystemPage, error) {
	obj := &SpacedRepetitionSystemPage{}
	err := c.request("GET", "/v2/spaced_repetition_systems", params, nil, obj)
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

type SpacedRepetitionSystem struct {
	Object
	Data *SpacedRepetitionSystemData `json:"data"`
}

type SpacedRepetitionSystemData struct {
	BurningStagePosition   int                                   `json:"burning_stage_position"`
	CreatedAt              time.Time                             `json:"created_at"`
	Description            string                                `json:"description"`
	Name                   string                                `json:"name"`
	PassingStagePosition   int                                   `json:"passing_stage_position"`
	Stages                 []*SpacedRepetitionSystemStagedObject `json:"stages"`
	StartingStagePosition  int                                   `json:"starting_stage_position"`
	UnlockingStagePosition int                                   `json:"unlocking_stage_position"`
}

type SpacedRepetitionSystemStagedObject struct {
	Interval     *int    `json:"interval"`
	IntervalUnit *string `json:"interval_unit"`
	Position     int     `json:"position"`
}

type SpacedRepetitionSystemGetParams struct {
	Params
	ID *WKID
}

type SpacedRepetitionSystemListParams struct {
	ListParams
	Params

	IDs          []WKID
	UpdatedAfter *WKTime
}

func (p *SpacedRepetitionSystemListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}
	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Encode())
	}

	return values.Encode()
}

type SpacedRepetitionSystemPage struct {
	*PageObject
	Data []*SpacedRepetitionSystem `json:"data"`
}
