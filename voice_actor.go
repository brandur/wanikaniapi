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

func (c *Client) VoiceActorGet(params *VoiceActorGetParams) (*VoiceActor, error) {
	obj := &VoiceActor{}
	err := c.request("GET", "/v2/voice_actors/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

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

type VoiceActor struct {
	Object
	Data *VoiceActorData `json:"data"`
}

type VoiceActorData struct {
	Description string `json:"description"`
	Gender      string `json:"gender"`
	Name        string `json:"name"`
}

type VoiceActorGetParams struct {
	Params
	ID *WKID
}

type VoiceActorListParams struct {
	ListParams
	Params

	IDs          []WKID
	UpdatedAfter *WKTime
}

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

type VoiceActorPage struct {
	*PageObject
	Data []*VoiceActor `json:"data"`
}
