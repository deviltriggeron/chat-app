package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"chat-app/internal/model"
	"chat-app/internal/ports"
)

type roomRepo struct {
	db *sql.DB
}

func NewRoomRepo(db *sql.DB) ports.RoomRepo {
	return &roomRepo{
		db: db,
	}
}

func (r *roomRepo) Insert(ctx context.Context, room model.RoomModel) error {
	q :=
		`
		INSERT INTO room (
			id, name, owner_id, created_at
		)
		VALUES (
			$1, $2, $3, $4
		)
	`

	_, err := r.db.ExecContext(ctx, q, room.ID, room.Name, room.OwnerID, room.OwnerID)
	if err != nil {
		return fmt.Errorf("error insert in database: %v", err)
	}

	return nil
}

func (r *roomRepo) Select(ctx context.Context, id int) (*model.RoomModel, error) {
	var room model.RoomModel

	q := `
		SELECT id, name, owner_id, created_at
		FROM room
		WHERE id = $1
	`

	err := r.db.QueryRowContext(ctx, q, id).Scan(&room.ID, &room.Name, &room.OwnerID, &room.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error select from database: %v", err)
	}

	return &room, nil
}

func (r *roomRepo) SelectAll(ctx context.Context) ([]model.RoomModel, error) {
	var rooms []model.RoomModel

	q := `
		SELECT id, name, owner_id, created_at
		FROM message
	`

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("error select all from database: %v", err)
	}

	for rows.Next() {
		var room model.RoomModel

		if err := rows.Scan(&room.ID, &room.Name, &room.OwnerID, &room.CreatedAt); err != nil {
			log.Println(err)
		}

		rooms = append(rooms, room)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	return rooms, nil
}

func (r *roomRepo) Delete(ctx context.Context, id int) error {
	q := `
		DELETE *
		FROM room
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("error delete in database: %v", err)
	}

	return nil
}
