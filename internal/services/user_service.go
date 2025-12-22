package services

import "chat-app/internal/ports"

type UserSvc struct {
	repo ports.UserRepo
}

func NewUserSvc(repo ports.UserRepo) UserService {
	return UserSvc{
		repo: repo,
	}
}
