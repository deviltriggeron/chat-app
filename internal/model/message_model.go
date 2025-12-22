package model

import "time"

type MessageModel struct {
	ID        int       `db:"id"`
	RoomID    int       `db:"room_id"`
	AuthorID  int       `db:"author_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
