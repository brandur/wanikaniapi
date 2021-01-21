package wanikaniapi

import (
	"strconv"
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

// VoiceActorGet retrieves a specific voice_actor by its ID.
func (c *Client) VoiceActorGet(params *VoiceActorGetParams) (*VoiceActor, error) {
	obj := &VoiceActor{}
	err := c.request("GET", "/v2/voice_actors/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

// VoiceActorList returns a collection of all voice actors, ordered by
// ascending CreatedAt, 500 at a time.
func (c *Client) VoiceActorList(params *VoiceActorListParams) (*VoiceActorPage, error) {
	obj := &VoiceActorPage{}
	err := c.request("GET", "/v2/voice_actors", params, nil, obj)
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

// VoiceActor represents a voice actor used for vocabulary reading
// pronunciation audio.
type VoiceActor struct {
	Object
	Data *VoiceActorData `json:"data"`
}

// VoiceActorData contains core data of VoiceActor.
type VoiceActorData struct {
	Description string `json:"description"`
	Gender      string `json:"gender"`
	Name        string `json:"name"`
}

// VoiceActorGetParams are parameters for VoiceActorGet.
type VoiceActorGetParams struct {
	Params
	ID *WKID
}

// VoiceActorListParams are parameters for VoiceActorList.
type VoiceActorListParams struct {
	ListParams
	Params

	IDs          []WKID
	UpdatedAfter *WKTime
}

// EncodeToQuery encodes parametes to a query string.
func (p *VoiceActorListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}

	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Encode())
	}

	return values.Encode()
}

// VoiceActorPage represents a single page of VoiceActors.
type VoiceActorPage struct {
	PageObject
	Data []*VoiceActor `json:"data"`
}
