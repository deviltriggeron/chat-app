package repository

import "chat-app/internal/ports"

type roomRepo struct {
}

func NewRoomRepo() ports.RoomRepo {
	return roomRepo{}
}
