package worshipup

import "time"

// SongJSON is a struct representation of the SongJSON format
type SongJSON struct {
	Title        string `json:"title"`
	Body         string `json:"body"`
	Presentation string `json:"presentation,omitempty"`
	Author       string `json:"author,omitempty"`
	CCLI         int    `json:"ccli,omitempty"`
}

// SetOrder holds a date and order in a service
type SetOrder struct {
	Date  time.Time
	Songs []string
}

// UUID is a universal unique identifier
type UUID string
