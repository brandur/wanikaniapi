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

// StudyMaterialCreate creates a study material for a specific subject.
//
// The owner of the API key can only create one study material per subject.
func (c *Client) StudyMaterialCreate(params *StudyMaterialCreateParams) (*StudyMaterial, error) {
	wrapper := &studyMaterialCreateParamsWrapper{Params: params.Params, StudyMaterial: params}
	obj := &StudyMaterial{}
	err := c.request("POST", "/v2/study_materials", params, wrapper, obj)
	return obj, err
}

// StudyMaterialGet retrieves a specific study material by its ID.
func (c *Client) StudyMaterialGet(params *StudyMaterialGetParams) (*StudyMaterial, error) {
	obj := &StudyMaterial{}
	err := c.request("GET", "/v2/study_materials/"+strconv.Itoa(int(*params.ID)), params, nil, obj)
	return obj, err
}

// StudyMaterialList returns a collection of all study material, ordered by
// ascending CreatedAt, 500 at a time.
func (c *Client) StudyMaterialList(params *StudyMaterialListParams) (*StudyMaterialPage, error) {
	obj := &StudyMaterialPage{}
	err := c.request("GET", "/v2/study_materials", params, nil, obj)
	return obj, err
}

// StudyMaterialUpdate updates a study material for a specific ID.
func (c *Client) StudyMaterialUpdate(params *StudyMaterialUpdateParams) (*StudyMaterial, error) {
	wrapper := &studyMaterialUpdateParamsWrapper{Params: params.Params, StudyMaterial: params}
	obj := &StudyMaterial{}
	err := c.request("PUT", "/v2/study_materials/"+strconv.Itoa(int(*params.ID)), params, wrapper, obj)
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

// StudyMaterial is a store of user-specific notes and synonyms for a given
// subject. The records are created as soon as the user enters any study
// information.
type StudyMaterial struct {
	Object
	Data *StudyMaterialData `json:"data"`
}

// StudyMaterialData contains core data of StudyMaterial.
type StudyMaterialData struct {
	CreatedAt       time.Time    `json:"created_at"`
	Hidden          bool         `json:"hidden"`
	MeaningNote     *string      `json:"meaning_note"`
	MeaningSynonyms []string     `json:"meaning_synonyms"`
	ReadingNote     *string      `json:"reading_note"`
	SubjectID       WKID         `json:"subject_id"`
	SubjectType     WKObjectType `json:"subject_type"`
}

// StudyMaterialCreateParams are parameters for StudyMaterialCreate.
type StudyMaterialCreateParams struct {
	Params
	MeaningNote     *string  `json:"meaning_note,omitempty"`
	MeaningSynonyms []string `json:"meaning_synonyms,omitempty"`
	ReadingNote     *string  `json:"reading_note,omitempty"`
	SubjectID       *WKID    `json:"subject_id,omitempty"`
}

type studyMaterialCreateParamsWrapper struct {
	Params
	StudyMaterial *StudyMaterialCreateParams `json:"study_material"`
}

// StudyMaterialGetParams are parameters for StudyMaterialGet.
type StudyMaterialGetParams struct {
	Params
	ID *WKID
}

// StudyMaterialListParams are parameters for StudyMaterialList.
type StudyMaterialListParams struct {
	ListParams
	Params

	Hidden       *bool
	IDs          []WKID
	SubjectIDs   []WKID
	SubjectTypes []WKObjectType
	UpdatedAfter *WKTime
}

// EncodeToQuery encodes parametes to a query string.
func (p *StudyMaterialListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.Hidden != nil {
		values.Add("hidden", strconv.FormatBool(*p.Hidden))
	}

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}

	if p.SubjectIDs != nil {
		values.Add("subject_ids", joinIDs(p.SubjectIDs, ","))
	}

	if p.SubjectTypes != nil {
		values.Add("subject_types", joinObjectTypes(p.SubjectTypes, ","))
	}

	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Encode())
	}

	return values.Encode()
}

// StudyMaterialPage represents a single page of StudyMaterials.
type StudyMaterialPage struct {
	PageObject
	Data []*StudyMaterial `json:"data"`
}

// StudyMaterialUpdateParams are parameters for StudyMaterialUpdate.
type StudyMaterialUpdateParams struct {
	Params
	ID              *WKID    `json:"-"`
	MeaningNote     *string  `json:"meaning_note,omitempty"`
	MeaningSynonyms []string `json:"meaning_synonyms,omitempty"`
	ReadingNote     *string  `json:"reading_note,omitempty"`
}

type studyMaterialUpdateParamsWrapper struct {
	Params
	StudyMaterial *StudyMaterialUpdateParams `json:"study_material"`
}
