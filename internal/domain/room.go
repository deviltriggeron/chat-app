package domain

import "time"

type Room struct {
	ID        int
	Name      string
	OwnerID   int
	CreatedAt time.Time
}
