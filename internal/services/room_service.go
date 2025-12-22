package services

import "chat-app/internal/ports"

type RoomSvc struct {
	repo ports.RoomRepo
}

func NewRoomSvc(repo ports.RoomRepo) RoomService {
	return RoomSvc{
		repo: repo,
	}
}
