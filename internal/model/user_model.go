package model

import "time"

type UserModel struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}
