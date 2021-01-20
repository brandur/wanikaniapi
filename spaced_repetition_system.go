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

// SpacedRepetitionSystemGet retrieves a specific spaced repetition system by
// its ID.
func (c *Client) SpacedRepetitionSystemGet(params *SpacedRepetitionSystemGetParams) (*SpacedRepetitionSystem, error) {
	obj := &SpacedRepetitionSystem{}
	err := c.request("GET", "/v2/spaced_repetition_systems/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

// SpacedRepetitionSystemList returns a collection of all spaced repetition
// systems, ordered by ascending ID, 500 at a time.
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

// SpacedRepetitionSystem is an available spaced repetition system used for
// calculating SRSStage changes to Assignments and Reviews. Has relationship
// with Subjects.
type SpacedRepetitionSystem struct {
	Object
	Data *SpacedRepetitionSystemData `json:"data"`
}

// SpacedRepetitionSystemData contains core data of SpacedRepetitionSystem.
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

// SpacedRepetitionSystemStagedObject represents an spaced repetition stage.
type SpacedRepetitionSystemStagedObject struct {
	Interval     *int    `json:"interval"`
	IntervalUnit *string `json:"interval_unit"`
	Position     int     `json:"position"`
}

// SpacedRepetitionSystemGetParams are parameters for SpacedRepetitionSystemGet.
type SpacedRepetitionSystemGetParams struct {
	Params
	ID *WKID
}

// SpacedRepetitionSystemListParams are parameters for SpacedRepetitionSystemList.
type SpacedRepetitionSystemListParams struct {
	ListParams
	Params

	IDs          []WKID
	UpdatedAfter *WKTime
}

// EncodeToQuery encodes parametes to a query string.
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

// SpacedRepetitionSystemPage represents a single page of SpacedRepetitionSystems.
type SpacedRepetitionSystemPage struct {
	PageObject
	Data []*SpacedRepetitionSystem `json:"data"`
}
