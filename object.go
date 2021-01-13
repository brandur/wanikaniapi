package wanikaniapi

import (
	// "encoding/json"
	"time"
)

type ID int64

type Object struct {
	DataUpdatedAt time.Time  `json:"data_updated_at"`
	ID            ID         `json:"id"`
	ObjectType    ObjectType `json:"object"`
	URL           string     `json:"url"`
}

type ObjectType string
