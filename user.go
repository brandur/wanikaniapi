package wanikaniapi

import (
	"time"
)

const (
	ObjectTypeUSer = ObjectType("user")
)

type User struct {
	Object
	Data *UserData `json:"data"`
}

type UserData struct {
	CurrentVacationStartedAt *time.Time             `json:"current_vacation_started_at"`
	ID                       string                 `json:"id"`
	Level                    int                    `json:"level"`
	Preferences              map[string]interface{} `json:"preferences"`
	ProfileURL               string                 `json:"profile_url"`
	StartedAt                time.Time              `json:"started_at"`
	Subscription             struct {
		Active          bool       `json:"active"`
		MaxLevelGranted int        `json:"max_level_granted"`
		PeriodEndsAt    *time.Time `json:"period_ends_at"`
		Type            string     `json:"type"`
	} `json:"subscription"`
	Username string `json:"username"`
}

type UserGetParams struct {
}

func (c *Client) UserGet(params *UserGetParams) (*User, error) {
	obj := &User{}
	err := c.request("GET", "/v2/user", "", obj)
	return obj, err
}
