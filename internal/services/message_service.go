package services

import "chat-app/internal/ports"

type messageSvc struct {
	repo ports.MessageRepo
}

func NewMsgSvc(repo ports.MessageRepo) MessageService {
	return &messageSvc{
		repo: repo,
	}
}
