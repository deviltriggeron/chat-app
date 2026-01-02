package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"chat-app/internal/model"
	"chat-app/internal/ports"
)

type msgRepo struct {
	db *sql.DB
}

func NewMsgRepo(db *sql.DB) ports.MessageRepo {
	return &msgRepo{
		db: db,
	}
}

func (m *msgRepo) Insert(ctx context.Context, msg model.MessageModel) error {
	q :=
		`
		INSERT INTO message (
			id, room_id, author_id, content, created_at
		)
		VALUES (
			$1, $2, $3, $4, $5
		)
	`

	_, err := m.db.ExecContext(ctx, q, msg.ID, msg.RoomID, msg.AuthorID, msg.Content, msg.CreatedAt)
	if err != nil {
		return fmt.Errorf("error insert in database: %v", err)
	}

	return nil
}

func (m *msgRepo) Select(ctx context.Context, id int) (*model.MessageModel, error) {
	var msg model.MessageModel

	q :=
		`
		SELECT id, room_id, author_id, content, created_at
		FROM message
		WHERE id = $1
	`

	err := m.db.QueryRowContext(ctx, q, id).Scan(&msg.ID, &msg.RoomID, &msg.AuthorID, &msg.Content, &msg.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error select from database: %v", err)
	}

	return &msg, nil
}

func (m *msgRepo) SelectAll(ctx context.Context) ([]model.MessageModel, error) {
	var messages []model.MessageModel

	q :=
		`
		SELECT id, room_id, author_id, content, created_at
		FROM message
	`

	rows, err := m.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("error select all from database: %v", err)
	}

	for rows.Next() {
		var msg model.MessageModel
		if err := rows.Scan(&msg.ID, &msg.RoomID, &msg.AuthorID, &msg.Content, &msg.CreatedAt); err != nil {
			log.Println(err)
		}
		messages = append(messages, msg)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	return messages, nil
}

func (m *msgRepo) Delete(ctx context.Context, id int) error {
	q :=
		`
		DELETE *
		FROM message
		WHERE id = $1
	`

	_, err := m.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("error delete in database: %v", err)
	}

	return nil
}
