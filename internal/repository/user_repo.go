package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"chat-app/internal/model"
	"chat-app/internal/ports"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) ports.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Insert(ctx context.Context, user model.UserModel) error {
	q := `
		INSERT INTO user (
			id, username, created_at
		)
		VALUES (
			$1, $2, $3
		)
	`

	_, err := u.db.ExecContext(ctx, q, user.ID, user.Username, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("error insert in database :%v", err)
	}

	return nil
}

func (u *userRepo) Select(ctx context.Context, id int) (*model.UserModel, error) {
	var user model.UserModel

	q := `
		SELECT id, username, created_at
		FROM user
		WHERE id = $1
	`

	err := u.db.QueryRowContext(ctx, q, id).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error select from database: %v", err)
	}

	return nil, nil
}

func (u *userRepo) SelectAll(ctx context.Context) ([]model.UserModel, error) {
	var users []model.UserModel

	q := `
		SELECT id, username, created_at
		FROM user
	`

	rows, err := u.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("error select from database: %v", err)
	}

	for rows.Next() {
		var user model.UserModel

		err := rows.Scan(&user.ID, &user.Username, &user.CreatedAt)
		if err != nil {
			log.Println(err)
		}

		users = append(users, user)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	return users, nil
}

func (u *userRepo) Delete(ctx context.Context, id int) error {
	q := `
		DELETE *
		FROM user
		WHERE id = $1
	`

	_, err := u.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("error delete in database: %v", err)
	}

	return nil
}
