package models

import (
	"time"
)

// TwitPost ...
type TwitPost struct {
	ID       string `json:"id,omitempty"`
	Text    string  `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
}
