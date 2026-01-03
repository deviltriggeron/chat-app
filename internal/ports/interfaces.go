package ports

import (
	"chat-app/internal/model"
	"context"
)

type MessageRepo interface {
	Insert(ctx context.Context, msg model.MessageModel) error
	Select(ctx context.Context, id int) (*model.MessageModel, error)
	SelectAll(ctx context.Context) ([]model.MessageModel, error)
	Delete(ctx context.Context, id int) error
}

type RoomRepo interface {
	Insert(ctx context.Context, room model.RoomModel) error
	Select(ctx context.Context, id int) (*model.RoomModel, error)
	SelectAll(ctx context.Context) ([]model.RoomModel, error)
	Delete(ctx context.Context, id int) error
}

type UserRepo interface {
	Insert(ctx context.Context, room model.UserModel) error
	Select(ctx context.Context, id int) (*model.UserModel, error)
	SelectAll(ctx context.Context) ([]model.UserModel, error)
	Delete(ctx context.Context, id int) error
}
