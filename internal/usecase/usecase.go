package usecase

import (
	"chat-app/internal/services"
)

type Usecase struct {
	message services.MessageService
	room    services.RoomService
	user    services.UserService
}

func NewUsecase(msg services.MessageService, room services.RoomService, user services.UserService) Usecase {
	return Usecase{
		message: msg,
		room:    room,
		user:    user,
	}
}
