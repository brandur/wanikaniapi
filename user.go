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

// UserGet returns a summary of user information.
func (c *Client) UserGet(params *UserGetParams) (*User, error) {
	obj := &User{}
	err := c.request("GET", "/v2/user", params, nil, obj)
	return obj, err
}

// UserUpdate returns an updated summary of user information.
func (c *Client) UserUpdate(params *UserUpdateParams) (*User, error) {
	wrapper := &userUpdateParamsWrapper{Params: params.Params, User: params}
	obj := &User{}
	err := c.request("PUT", "/v2/user", params, wrapper, obj)
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

// User returns basic information for the user making the API request,
// identified by their API token.
type User struct {
	Object
	Data *UserData `json:"data"`
}

// UserData contains core data of User.
type UserData struct {
	CurrentVacationStartedAt *time.Time        `json:"current_vacation_started_at"`
	WKID                     string            `json:"id"`
	Level                    int               `json:"level"`
	Preferences              *UserPreferences  `json:"preferences"`
	ProfileURL               string            `json:"profile_url"`
	StartedAt                time.Time         `json:"started_at"`
	Subscription             *UserSubscription `json:"subscription"`
	Username                 string            `json:"username"`
}

// UserGetParams are parameters for UserGet.
type UserGetParams struct {
	Params
}

// UserPreferences are preferences for a user.
type UserPreferences struct {
	DefaultVoiceActorID        WKID   `json:"default_voice_actor_id"`
	LessonsAutoplayAudio       bool   `json:"lessons_autoplay_audio"`
	LessonsBatchSize           int    `json:"lessons_batch_size"`
	LessonsPresentationOrder   string `json:"lessons_presentation_order"`
	ReviewsAutoplayAudio       bool   `json:"reviews_autoplay_audio"`
	ReviewsDisplaySRSIndicator bool   `json:"reviews_display_srs_indicator"`
}

// UserSubscription represents a subscription for a user.
type UserSubscription struct {
	Active          bool       `json:"active"`
	MaxLevelGranted int        `json:"max_level_granted"`
	PeriodEndsAt    *time.Time `json:"period_ends_at"`
	Type            string     `json:"type"`
}

// UserUpdateParams are parameters for UserUpdate.
type UserUpdateParams struct {
	Params
	Preferences *UserUpdatePreferencesParams `json:"preferences,omitempty"`
}

// UserUpdatePreferencesParams are update parameters for user preferences.
type UserUpdatePreferencesParams struct {
	DefaultVoiceActorID        *WKID   `json:"default_voice_actor_id,omitempty"`
	LessonsAutoplayAudio       *bool   `json:"lessons_autoplay_audio,omitempty"`
	LessonsBatchSize           *int    `json:"lessons_batch_size,omitempty"`
	LessonsPresentationOrder   *string `json:"lessons_presentation_order,omitempty"`
	ReviewsAutoplayAudio       *bool   `json:"reviews_autoplay_audio,omitempty"`
	ReviewsDisplaySRSIndicator *bool   `json:"reviews_display_srs_indicator,omitempty"`
}

type userUpdateParamsWrapper struct {
	Params
	User *UserUpdateParams `json:"user"`
}
