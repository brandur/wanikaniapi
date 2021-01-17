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

func (c *Client) StudyMaterialCreate(params *StudyMaterialCreateParams) (*StudyMaterial, error) {
	wrapper := &studyMaterialCreateParamsWrapper{Params: params}
	obj := &StudyMaterial{}
	err := c.request("POST", "/v2/study_materials", "", wrapper, obj)
	return obj, err
}

func (c *Client) StudyMaterialGet(params *StudyMaterialGetParams) (*StudyMaterial, error) {
	obj := &StudyMaterial{}
	err := c.request("GET", "/v2/study_materials/"+strconv.Itoa(int(*params.ID)), "", nil, obj)
	return obj, err
}

func (c *Client) StudyMaterialList(params *StudyMaterialListParams) (*StudyMaterialPage, error) {
	obj := &StudyMaterialPage{}
	err := c.request("GET", "/v2/study_materials", params.EncodeToQuery(), nil, obj)
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

type StudyMaterial struct {
	Object
	Data *StudyMaterialData `json:"data"`
}

type StudyMaterialData struct {
	CreatedAt       time.Time  `json:"created_at"`
	Hidden          bool       `json:"hidden"`
	MeaningNote     *string    `json:"meaning_note"`
	MeaningSynonyms []string   `json:"meaning_synonyms"`
	ReadingNote     *string    `json:"reading_note"`
	SubjectID       ID         `json:"subject_id"`
	SubjectType     ObjectType `json:"subject_type"`
}

type StudyMaterialCreateParams struct {
	*ListParams
	MeaningNote     *string  `json:"meaning_note,omitempty"`
	MeaningSynonyms []string `json:"meaning_synonyms,omitempty"`
	ReadingNote     *string  `json:"reading_note,omitempty"`
	SubjectID       *ID      `json:"subject_id,omitempty"`
}

type studyMaterialCreateParamsWrapper struct {
	Params *StudyMaterialCreateParams `json:"study_material"`
}

type StudyMaterialGetParams struct {
	ID *ID
}

type StudyMaterialListParams struct {
	*ListParams
	Hidden       *bool
	IDs          []ID
	SubjectIDs   []ID
	SubjectTypes []ObjectType
	UpdatedAfter *time.Time
}

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
		values.Add("updated_after", p.UpdatedAfter.Format(time.RFC3339))
	}

	return values.Encode()
}

type StudyMaterialPage struct {
	*PageObject
	Data []*StudyMaterial `json:"data"`
}
