-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS message (
    id INTEGER,
    room_id INTEGER,
    author_id INTEGER,
    content INTEGER,
    created_at DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS message;
-- +goose StatementEnd