package domain

import "time"

type Message struct {
	ID        int
	RoomID    int
	AuthorID  int
	Content   string
	CreatedAt time.Time
}
