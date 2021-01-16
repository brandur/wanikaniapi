package wanikaniapi

import (
	"encoding/json"
	"fmt"
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

func (c *Client) SubjectGet(params *SubjectGetParams) (*Subject, error) {
	obj := &Subject{}
	err := c.request("GET", "/v2/subjects/"+strconv.FormatInt(int64(*params.ID), 10), "", nil, obj)
	return obj, err
}

func (c *Client) SubjectList(params *SubjectListParams) (*SubjectPage, error) {
	obj := &SubjectPage{}
	err := c.request("GET", "/v2/subjects", params.EncodeToQuery(), nil, obj)
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

type Subject struct {
	Object

	KanjiData      *SubjectKanjiData
	RadicalData    *SubjectRadicalData
	VocabularyData *SubjectVocabularyData
}

func (s *Subject) UnmarshalJSON(data []byte) error {
	type subject Subject
	var v subject
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = Subject(v)

	var objMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return err
	}

	switch s.Object.ObjectType {
	case ObjectTypeKanji:
		if _, ok := objMap["data"]; ok {
			s.KanjiData = &SubjectKanjiData{}
			if err := json.Unmarshal(objMap["data"], s.KanjiData); err != nil {
				return fmt.Errorf("decoding kanji from subject: %w", err)
			}
		}

	case ObjectTypeRadical:
		if _, ok := objMap["data"]; ok {
			s.RadicalData = &SubjectRadicalData{}
			if err := json.Unmarshal(objMap["data"], s.RadicalData); err != nil {
				return fmt.Errorf("decoding radical from subject: %w", err)
			}
		}

	case ObjectTypeVocabulary:
		if _, ok := objMap["data"]; ok {
			s.VocabularyData = &SubjectVocabularyData{}
			if err := json.Unmarshal(objMap["data"], s.VocabularyData); err != nil {
				return fmt.Errorf("decoding vocabulary from subject: %w", err)
			}
		}
	}

	return nil
}

type SubjectCommonData struct {
	AuxiliaryMeanings        []*SubjectAuxiliaryMeaningObject `json:"auxiliary_meanings"`
	CreatedAt                time.Time                        `json:"created_at"`
	DocumentURL              string                           `json:"document_url"`
	HiddenAt                 *time.Time                       `json:"hidden_at"`
	Level                    int                              `json:"level"`
	LessonPosition           int                              `json:"lesson_position"`
	Meanings                 []*SubjectMeaningObject          `json:"meanings"`
	Slug                     string                           `json:"slug"`
	SpacedRepetitionSystemID ID                               `json:"spaced_repetition_system_id"`
}

type SubjectKanjiData struct {
	SubjectCommonData

	AmalgamationSubjectIDs    []ID                   `json:"amalgamation_subject_ids"`
	Characters                string                 `json:"characters"`
	ComponentSubjectIDs       []ID                   `json:"component_subject_ids"`
	MeaningHint               *string                `json:"meaning_hint"`
	ReadingHint               *string                `json:"reading_hint"`
	ReadingMnemonic           string                 `json:"mnemonic_hint"`
	Readings                  []*SubjectKanjiReading `json:"readings"`
	VisuallySimilarSubjectIDs []ID                   `json:"visually_similar_subject_ids"`
}

type SubjectKanjiReading struct {
	AcceptedAnswer bool                    `json:"accepted_answer"`
	Primary        bool                    `json:"primary"`
	Reading        string                  `json:"reading"`
	Type           SubjectKanjiReadingType `json:"type"`
}

type SubjectKanjiReadingType string

const (
	SubjectKanjiReadingTypeKunyomi SubjectKanjiReadingType = "kunyomi"
	SubjectKanjiReadingTypeNanori  SubjectKanjiReadingType = "nanori"
	SubjectKanjiReadingTypeOnyomi  SubjectKanjiReadingType = "onyomi"
)

type SubjectRadicalCharacterImage struct {
	ContentType string                                                  `json:"content_type"`
	Metadata    map[SubjectRadicalCharacterImageMetadataKey]interface{} `json:"metadata"`
	URL         string                                                  `json:"url"`
}

type SubjectRadicalCharacterImageMetadataKey string

const (
	SubjectRadicalCharacterImageMetadataKeyColor        SubjectRadicalCharacterImageMetadataKey = "color"
	SubjectRadicalCharacterImageMetadataKeyDimensions   SubjectRadicalCharacterImageMetadataKey = "dimensions"
	SubjectRadicalCharacterImageMetadataKeyInlineStyles SubjectRadicalCharacterImageMetadataKey = "inline_styles"
	SubjectRadicalCharacterImageMetadataKeyStyleName    SubjectRadicalCharacterImageMetadataKey = "style_name"
)

type SubjectRadicalData struct {
	SubjectCommonData

	AmalgamationSubjectIDs []ID                            `json:"amalgamation_subject_ids"`
	CharacterImages        []*SubjectRadicalCharacterImage `json:"character_images"`
	Characters             *string                         `json:"characters"`
}

type SubjectVocabularyContextSentence struct {
	EN string `json:"en"`
	JA string `json:"ja"`
}

type SubjectVocabularyData struct {
	SubjectCommonData

	Characters           string                                 `json:"characters"`
	ComponentSubjectIDs  []ID                                   `json:"component_subject_ids"`
	ContextSentences     []*SubjectVocabularyContextSentence    `json:"context_sentences"`
	MeaningMnemonic      string                                 `json:"meaning_mnenomic"`
	PartsOfSpeech        []string                               `json:"parts_of_speech"`
	PronounciationAudios []SubjectVocabularyPronounciationAudio `json:"pronounciation_audios"`
	Readings             []*SubjectVocabularyReading            `json:"subject_vocabulary_reading"`
}

type SubjectVocabularyPronounciationAudio struct {
	ContentType string                                                          `json:"content_type"`
	Metadata    map[SubjectVocabularyPronounciationAudioMetadataKey]interface{} `json:"metadata"`
	URL         string                                                          `json:"url"`
}

type SubjectVocabularyPronounciationAudioMetadataKey string

const (
	SubjectVocabularyPronounciationAudioMetadataKeyGender           SubjectVocabularyPronounciationAudioMetadataKey = "gender"
	SubjectVocabularyPronounciationAudioMetadataKeyPronounciation   SubjectVocabularyPronounciationAudioMetadataKey = "pronounciation"
	SubjectVocabularyPronounciationAudioMetadataKeySourceID         SubjectVocabularyPronounciationAudioMetadataKey = "source_id"
	SubjectVocabularyPronounciationAudioMetadataKeyVoiceActorID     SubjectVocabularyPronounciationAudioMetadataKey = "voice_actor_id"
	SubjectVocabularyPronounciationAudioMetadataKeyVoiceActorName   SubjectVocabularyPronounciationAudioMetadataKey = "voice_actor_name"
	SubjectVocabularyPronounciationAudioMetadataKeyVoiceDescription SubjectVocabularyPronounciationAudioMetadataKey = "voice_description"
)

type SubjectVocabularyReading struct {
	AcceptedAnswer bool   `json:"accepted_answer"`
	Primary        bool   `json:"primary"`
	Reading        string `json:"reading"`
}

type SubjectGetParams struct {
	ID *ID
}

type SubjectListParams struct {
	*ListParams
	IDs          []ID
	Hidden       *bool
	Levels       []int
	Slugs        []string
	Types        []string
	UpdatedAfter *time.Time
}

func (p *SubjectListParams) EncodeToQuery() string {
	values := p.encodeToURLValues()

	if p.IDs != nil {
		values.Add("ids", joinIDs(p.IDs, ","))
	}

	if p.Hidden != nil {
		values.Add("hidden", strconv.FormatBool(*p.Hidden))
	}

	if p.Levels != nil {
		values.Add("levels", joinInts(p.Levels, ","))
	}

	if p.Slugs != nil {
		values.Add("slugs", joinStrings(p.Slugs, ","))
	}

	if p.Types != nil {
		values.Add("types", joinStrings(p.Types, ","))
	}

	if p.UpdatedAfter != nil {
		values.Add("updated_after", p.UpdatedAfter.Format(time.RFC3339))
	}

	return values.Encode()
}

type SubjectAuxiliaryMeaningObject struct {
	Meaning string                            `json:"meaning"`
	Type    SubjectAuxiliaryMeaningObjectType `json:"type"`
}

type SubjectAuxiliaryMeaningObjectType string

const (
	SubjectAuxiliaryMeaningObjectTypeBlacklist SubjectAuxiliaryMeaningObjectType = "blacklist"
	SubjectAuxiliaryMeaningObjectTypeWhitelist SubjectAuxiliaryMeaningObjectType = "whitelist"
)

type SubjectMeaningObject struct {
	AcceptedAnswer bool   `json:"accepted_answer"`
	Meaning        string `json:"meaning"`
	Primary        bool   `json:"primary"`
}

type SubjectPage struct {
	*PageObject
	Data []*Subject `json:"data"`
}
