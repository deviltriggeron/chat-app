package repository

import "chat-app/internal/ports"

type msgRepo struct {
}

func NewMsgRepo() ports.MessageRepo {
	return msgRepo{}
}
