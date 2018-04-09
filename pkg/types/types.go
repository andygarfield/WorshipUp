package types

import (
	"time"

	"github.com/andygarfield/worshipup/pkg/utils"
)

// SongJSON is a struct representation of the SongJSON format
type SongJSON struct {
	ID           uint64  `json:"ID"`
	Title        string  `json:"title"`
	Body         string  `json:"body"`
	Presentation *string `json:"presentation,omitempty"`
	Author       *string `json:"author,omitempty"`
	CCLI         *int32  `json:"ccli,omitempty"`
}

// SetList holds a date and order in a service
type SetList struct {
	ID    uint64
	Date  time.Time
	Songs *[]*uint64
}

// Scrub cleans user song input
func (s *SongJSON) Scrub() error {
	var err error
	s.Title, err = utils.ScrubUserTitle(s.Title)
	if err != nil {
		return err
	}
	s.Body, err = utils.ScrubUserBody(s.Body)
	if err != nil {
		return err
	}
	return nil
}
