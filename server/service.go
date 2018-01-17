package main

import (
	"time"
)

type serviceOrder struct {
	Date  time.Time
	Songs []SongJSON
}
