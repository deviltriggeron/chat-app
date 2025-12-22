package repository

import "chat-app/internal/ports"

type userRepo struct {
}

func NewUserRepo() ports.UserRepo {
	return userRepo{}
}
