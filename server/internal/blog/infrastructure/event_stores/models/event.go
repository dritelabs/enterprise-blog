package models

import "time"

type Event struct {
	CreatedAt time.Time
	Payload   any
	Type      string
}
