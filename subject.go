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

// SubjectGet retrieves a specific subject by its ID. The structure of the
// response depends on the subject type.
func (c *Client) SubjectGet(params *SubjectGetParams) (*Subject, error) {
	obj := &Subject{}
	err := c.request("GET", "/v2/subjects/"+strconv.FormatInt(int64(*params.ID), 10), params, nil, obj)
	return obj, err
}

// SubjectList returns a collection of all subjects, ordered by ascending
// CreatedAt, 1000 at a time.
func (c *Client) SubjectList(params *SubjectListParams) (*SubjectPage, error) {
	obj := &SubjectPage{}
	err := c.request("GET", "/v2/subjects", params, nil, obj)
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

// Subject represents radicals, kanji, and vocabulary that are learned through
// lessons and reviews. They contain basic dictionary information, such as
// meanings and/or readings, and information about their relationship to other
// items with WaniKani, like their level.
type Subject struct {
	Object

	// KanjiData is data on a kanji subject. Populated only if he subject is a
	// kanji.
	KanjiData *SubjectKanjiData

	// RadicalData is data on a radical subject. Populated only if he subject
	// is a radical.
	RadicalData *SubjectRadicalData

	// VocabularyData is data on a vocabulary subject. Populated only if he
	// subject is a vocabulary.
	VocabularyData *SubjectVocabularyData
}

// UnmarshalJSON is a custom JSON unmarshaling function for Subject.
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

// SubjectCommonData is common data available on all subject types regardless
// of whether they're kanji, radical, or vocabulary.
type SubjectCommonData struct {
	AuxiliaryMeanings        []*SubjectAuxiliaryMeaningObject `json:"auxiliary_meanings"`
	CreatedAt                time.Time                        `json:"created_at"`
	DocumentURL              string                           `json:"document_url"`
	HiddenAt                 *time.Time                       `json:"hidden_at"`
	Level                    int                              `json:"level"`
	LessonPosition           int                              `json:"lesson_position"`
	Meanings                 []*SubjectMeaningObject          `json:"meanings"`
	Slug                     string                           `json:"slug"`
	SpacedRepetitionSystemID WKID                             `json:"spaced_repetition_system_id"`
}

// SubjectKanjiData is data on a kanji subject.
type SubjectKanjiData struct {
	SubjectCommonData

	AmalgamationSubjectIDs    []WKID                 `json:"amalgamation_subject_ids"`
	Characters                string                 `json:"characters"`
	ComponentSubjectIDs       []WKID                 `json:"component_subject_ids"`
	MeaningHint               *string                `json:"meaning_hint"`
	ReadingHint               *string                `json:"reading_hint"`
	ReadingMnemonic           string                 `json:"mnemonic_hint"`
	Readings                  []*SubjectKanjiReading `json:"readings"`
	VisuallySimilarSubjectIDs []WKID                 `json:"visually_similar_subject_ids"`
}

// SubjectKanjiReading is reading data on a kanji subject.
type SubjectKanjiReading struct {
	AcceptedAnswer bool                    `json:"accepted_answer"`
	Primary        bool                    `json:"primary"`
	Reading        string                  `json:"reading"`
	Type           SubjectKanjiReadingType `json:"type"`
}

// SubjectKanjiReadingType is the type of a kanji reading.
type SubjectKanjiReadingType string

// All possible types of kanji readings.
const (
	SubjectKanjiReadingTypeKunyomi SubjectKanjiReadingType = "kunyomi"
	SubjectKanjiReadingTypeNanori  SubjectKanjiReadingType = "nanori"
	SubjectKanjiReadingTypeOnyomi  SubjectKanjiReadingType = "onyomi"
)

// SubjectRadicalCharacterImage represents an image for a radical. Unlike kanji
// or vocabulary, radicals cannot always be represented by a unicode glyph.
type SubjectRadicalCharacterImage struct {
	ContentType string                                                  `json:"content_type"`
	Metadata    map[SubjectRadicalCharacterImageMetadataKey]interface{} `json:"metadata"`
	URL         string                                                  `json:"url"`
}

// SubjectRadicalCharacterImageMetadataKey is a key for character image metadata.
type SubjectRadicalCharacterImageMetadataKey string

// All possible values of radical character image metadata keys.
const (
	SubjectRadicalCharacterImageMetadataKeyColor        SubjectRadicalCharacterImageMetadataKey = "color"
	SubjectRadicalCharacterImageMetadataKeyDimensions   SubjectRadicalCharacterImageMetadataKey = "dimensions"
	SubjectRadicalCharacterImageMetadataKeyInlineStyles SubjectRadicalCharacterImageMetadataKey = "inline_styles"
	SubjectRadicalCharacterImageMetadataKeyStyleName    SubjectRadicalCharacterImageMetadataKey = "style_name"
)

// SubjectRadicalData is data on a radical subject.
type SubjectRadicalData struct {
	SubjectCommonData

	AmalgamationSubjectIDs []WKID                          `json:"amalgamation_subject_ids"`
	CharacterImages        []*SubjectRadicalCharacterImage `json:"character_images"`
	Characters             *string                         `json:"characters"`
}

// SubjectVocabularyContextSentence represents a vocabulary context sentence.
type SubjectVocabularyContextSentence struct {
	// EN is the English translation of the sentence.
	EN string `json:"en"`

	// JA is a Japanese context context sentence.
	JA string `json:"ja"`
}

// SubjectVocabularyData is data on a vocabulary subject.
type SubjectVocabularyData struct {
	SubjectCommonData

	Characters           string                                  `json:"characters"`
	ComponentSubjectIDs  []WKID                                  `json:"component_subject_ids"`
	ContextSentences     []*SubjectVocabularyContextSentence     `json:"context_sentences"`
	MeaningMnemonic      string                                  `json:"meaning_mnenomic"`
	PartsOfSpeech        []string                                `json:"parts_of_speech"`
	PronounciationAudios []*SubjectVocabularyPronounciationAudio `json:"pronounciation_audios"`
	Readings             []*SubjectVocabularyReading             `json:"subject_vocabulary_reading"`
}

// SubjectVocabularyPronounciationAudio represets an audio object for
// vocabulary pronounciation.
type SubjectVocabularyPronounciationAudio struct {
	ContentType string                                                          `json:"content_type"`
	Metadata    map[SubjectVocabularyPronounciationAudioMetadataKey]interface{} `json:"metadata"`
	URL         string                                                          `json:"url"`
}

// SubjectVocabularyPronounciationAudioMetadataKey is a key for pronounciation
// audio metadata.
type SubjectVocabularyPronounciationAudioMetadataKey string

// All possible values of vocabulary pronounciation audio metadata keys.
const (
	SubjectVocabularyPronounciationAudioMetadataKeyGender           SubjectVocabularyPronounciationAudioMetadataKey = "gender"
	SubjectVocabularyPronounciationAudioMetadataKeyPronounciation   SubjectVocabularyPronounciationAudioMetadataKey = "pronounciation"
	SubjectVocabularyPronounciationAudioMetadataKeySourceID         SubjectVocabularyPronounciationAudioMetadataKey = "source_id"
	SubjectVocabularyPronounciationAudioMetadataKeyVoiceActorID     SubjectVocabularyPronounciationAudioMetadataKey = "voice_actor_id"
	SubjectVocabularyPronounciationAudioMetadataKeyVoiceActorName   SubjectVocabularyPronounciationAudioMetadataKey = "voice_actor_name"
	SubjectVocabularyPronounciationAudioMetadataKeyVoiceDescription SubjectVocabularyPronounciationAudioMetadataKey = "voice_description"
)

// SubjectVocabularyReading is reading data on a vocabulary subject.
type SubjectVocabularyReading struct {
	AcceptedAnswer bool   `json:"accepted_answer"`
	Primary        bool   `json:"primary"`
	Reading        string `json:"reading"`
}

// SubjectGetParams are parameters for SubjectGet.
type SubjectGetParams struct {
	Params
	ID *WKID
}

// SubjectListParams are parameters for SubjectList.
type SubjectListParams struct {
	ListParams
	Params

	IDs          []WKID
	Hidden       *bool
	Levels       []int
	Slugs        []string
	Types        []string
	UpdatedAfter *WKTime
}

// EncodeToQuery encodes parametes to a query string.
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
		values.Add("updated_after", p.UpdatedAfter.Encode())
	}

	return values.Encode()
}

// SubjectAuxiliaryMeaningObject represents an auxiliary meaning for a subject.
type SubjectAuxiliaryMeaningObject struct {
	Meaning string                            `json:"meaning"`
	Type    SubjectAuxiliaryMeaningObjectType `json:"type"`
}

// SubjectAuxiliaryMeaningObjectType is the type of object of an auxiliary meaning.
type SubjectAuxiliaryMeaningObjectType string

// All possible values of auxiliary meaning object type.
const (
	SubjectAuxiliaryMeaningObjectTypeBlacklist SubjectAuxiliaryMeaningObjectType = "blacklist"
	SubjectAuxiliaryMeaningObjectTypeWhitelist SubjectAuxiliaryMeaningObjectType = "whitelist"
)

// SubjectMeaningObject represents a meaning for a subject.
type SubjectMeaningObject struct {
	AcceptedAnswer bool   `json:"accepted_answer"`
	Meaning        string `json:"meaning"`
	Primary        bool   `json:"primary"`
}

// SubjectPage represents a single page of Subjects.
type SubjectPage struct {
	PageObject
	Data []*Subject `json:"data"`
}
