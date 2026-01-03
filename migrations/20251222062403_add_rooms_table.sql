-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS room (
    id INTEGER,
    name TEXT,
    owner_id INTEGER,
    created_at DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS room;
-- +goose StatementEnd

type RoomModel struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	OwnerID   int       `db:"owner_id"`
	CreatedAt time.Time `db:"created_at"`
}