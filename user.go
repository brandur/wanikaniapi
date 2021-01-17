package wanikaniapi

import (
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

func (c *Client) UserGet(params *UserGetParams) (*User, error) {
	obj := &User{}
	err := c.request("GET", "/v2/user", "", nil, obj)
	return obj, err
}

func (c *Client) UserUpdate(params *UserUpdateParams) (*User, error) {
	wrapper := &userUpdateParamsWrapper{Params: params}
	obj := &User{}
	err := c.request("PUT", "/v2/user", "", wrapper, obj)
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

type User struct {
	Object
	Data *UserData `json:"data"`
}

type UserData struct {
	CurrentVacationStartedAt *time.Time        `json:"current_vacation_started_at"`
	ID                       string            `json:"id"`
	Level                    int               `json:"level"`
	Preferences              *UserPreferences  `json:"preferences"`
	ProfileURL               string            `json:"profile_url"`
	StartedAt                time.Time         `json:"started_at"`
	Subscription             *UserSubscription `json:"subscription"`
	Username                 string            `json:"username"`
}

type UserGetParams struct {
}

type UserPreferences struct {
	DefaultVoiceActorID        ID     `json:"default_voice_actor_id"`
	LessonsAutoplayAudio       bool   `json:"lessons_autoplay_audio"`
	LessonsBatchSize           int    `json:"lessons_batch_size"`
	LessonsPresentationOrder   string `json:"lessons_presentation_order"`
	ReviewsAutoplayAudio       bool   `json:"reviews_autoplay_audio"`
	ReviewsDisplaySRSIndicator bool   `json:"reviews_display_srs_indicator"`
}

type UserSubscription struct {
	Active          bool       `json:"active"`
	MaxLevelGranted int        `json:"max_level_granted"`
	PeriodEndsAt    *time.Time `json:"period_ends_at"`
	Type            string     `json:"type"`
}

type UserUpdateParams struct {
	Preferences *UserUpdatePreferencesParams `json:"preferences,omitempty"`
}

type UserUpdatePreferencesParams struct {
	DefaultVoiceActorID        *ID     `json:"default_voice_actor_id,omitempty"`
	LessonsAutoplayAudio       *bool   `json:"lessons_autoplay_audio,omitempty"`
	LessonsBatchSize           *int    `json:"lessons_batch_size,omitempty"`
	LessonsPresentationOrder   *string `json:"lessons_presentation_order,omitempty"`
	ReviewsAutoplayAudio       *bool   `json:"reviews_autoplay_audio,omitempty"`
	ReviewsDisplaySRSIndicator *bool   `json:"reviews_display_srs_indicator,omitempty"`
}

type userUpdateParamsWrapper struct {
	Params *UserUpdateParams `json:"user"`
}
